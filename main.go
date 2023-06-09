package main

import (
	"context"
	"fmt"
	"github.com/application-research/delta-metrics-rest/api"
	"github.com/application-research/delta-metrics-rest/dao"
	"github.com/application-research/delta-metrics-rest/model"
	"github.com/droundy/goopt"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	explru "github.com/paskal/golang-lru/simplelru"
	"github.com/spf13/viper"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const CacheSize = 1024 * 1024 * 1024 // 1GB
const CacheDuration = time.Hour * 4
const CachePurgeEveryDuration = time.Hour * 4

var (
	// BuildDate date string of when build was performed filled in by -X compile flag
	BuildDate string

	// LatestCommit date string of when build was performed filled in by -X compile flag
	LatestCommit string

	// BuildNumber date string of when build was performed filled in by -X compile flag
	BuildNumber string

	// BuiltOnIP date string of when build was performed filled in by -X compile flag
	BuiltOnIP string

	// BuiltOnOs date string of when build was performed filled in by -X compile flag
	BuiltOnOs string

	// RuntimeVer date string of when build was performed filled in by -X compile flag
	RuntimeVer string

	// OsSignal signal used to shutdown
	OsSignal chan os.Signal
)

// GinServer launch gin server
func GinServer() (err error) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api.ConfigGinRouter(router)
	router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server, the error is '%v'", err)
	}

	return
}

// @title Sample CRUD api for estuary db
// @version 1.0
// @description Sample CRUD api for estuary db
// @termsOfService

// @contact.name Me
// @contact.url http://me.com/terms.html
// @contact.email alvin@protocol.ai

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	OsSignal = make(chan os.Signal, 1)
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	dbHost, okHost := viper.Get("DB_HOST").(string)
	dbUser, okUser := viper.Get("DB_USER").(string)
	dbPass, okPass := viper.Get("DB_PASS").(string)
	dbName, okName := viper.Get("DB_NAME").(string)
	dbPort, okPort := viper.Get("DB_PORT").(string)
	if !okHost || !okUser || !okPass || !okName || !okPort {
		log.Fatalf("Error while reading database config")
	}

	// Define version information
	goopt.Version = fmt.Sprintf(
		`Application build information
				  Build date      : %s
				  Build number    : %s
				  Git commit      : %s
				  Runtime version : %s
				  Built on OS     : %s
				`, BuildDate, BuildNumber, LatestCommit, RuntimeVer, BuiltOnOs)
	goopt.Parse(nil)

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort

	// Define version information
	goopt.Version = fmt.Sprintf(
		`Application build information
  Build date      : %s
  Build number    : %s
  Git commit      : %s
  Runtime version : %s
  Built on OS     : %s
`, BuildDate, BuildNumber, LatestCommit, RuntimeVer, BuiltOnOs)
	goopt.Parse(nil)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(true)

	// Get generic database object sql.DB to use its functions
	sqlDB := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	dao.DB = db

	db.AutoMigrate(
		&model.ContentDealLogs{},
		&model.ContentDealProposalLogs{},
		&model.ContentDealProposalParametersLogs{},
		&model.ContentLogs{},
		&model.ContentMinerLogs{},
		&model.ContentWalletLogs{},
		&model.DeltaNodeGeoLocations{},
		&model.DeltaStartupLogs{},
		&model.InstanceMetaLogs{},
		&model.LogEvents{},
		&model.PieceCommitmentLogs{},
		&model.WalletLogs{},
	)

	dao.Logger = func(ctx context.Context, sql string) {
		fmt.Printf("SQL: %s\n", sql)
	}

	// cache
	dao.Cacher = explru.NewExpirableLRU(CacheSize, nil, CacheDuration, CachePurgeEveryDuration)

	// Initialize Refresh Views
	RefreshDBViews()

	go GinServer()
	LoopForever()
}

func RefreshDBViews() {

	s := gocron.NewScheduler(time.UTC)

	// Every starts the job immediately and then runs at the
	// specified interval
	jobRefreshGlobalStats, err := s.Every(4).Hours().Do(func() {
		fmt.Println("Refresh Stats Views")
		refreshViews, err := os.ReadFile("sql/views/refresh_mv_stats.sql")
		refreshViewsStr := string(refreshViews)
		if err != nil {
			log.Printf("Got error when reading refresh_mv_stats.sql, the error is '%v'", err)
		}
		if err := dao.DB.Exec(refreshViewsStr); err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Println(err)
	}

	jobRefreshAllTables, err := s.Every(4).Hours().Do(func() {
		fmt.Println("Refresh All Table Views")
		refreshViews, err := os.ReadFile("sql/views/refresh_all_tables.sql")
		refreshViewsStr := string(refreshViews)
		if err != nil {
			log.Println(err)
		}
		if err := dao.DB.Exec(refreshViewsStr); err != nil {
			log.Println(err)
		}
	})

	fmt.Println(jobRefreshGlobalStats)
	fmt.Println(jobRefreshAllTables)

	s.StartAsync()
}

// LoopForever on signal processing
func LoopForever() {
	fmt.Printf("Entering infinite loop\n")

	signal.Notify(OsSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	_ = <-OsSignal

	fmt.Printf("Exiting infinite loop received OsSignal\n")

}
