package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: content_miner_logs
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 2] miner                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] system_content_miner_id                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 7] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 8] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 9] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 55,    "content": 74,    "miner": "eZUwbMvFVdQoZqjpXohTQQJUW",    "node_info": "VKEmWOogdcixMZtPKIEItLOWX",    "requester_info": "QDsVNtMlcnNtHaPOtrHocJZQI",    "requesting_api_key": "oxFMmJQdJwRSCxfjFEOEMBPNF",    "system_content_miner_id": 77,    "created_at": "2042-06-30T04:08:40.293262949-04:00",    "updated_at": "2063-06-07T15:45:47.396911388-04:00",    "delta_node_uuid": "sDGcAFphRhhGDXdUNJSYGLjiL"}



*/

// ContentMinerLogs struct is a row record of the content_miner_logs table in the estuary database
type ContentMinerLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Content sql.NullInt64 `gorm:"column:content;type:INT8;"`
	//[ 2] miner                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Miner sql.NullString `gorm:"column:miner;type:TEXT;"`
	//[ 3] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo sql.NullString `gorm:"column:node_info;type:TEXT;"`
	//[ 4] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo sql.NullString `gorm:"column:requester_info;type:TEXT;"`
	//[ 5] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequestingAPIKey sql.NullString `gorm:"column:requesting_api_key;type:TEXT;"`
	//[ 6] system_content_miner_id                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemContentMinerID sql.NullInt64 `gorm:"column:system_content_miner_id;type:INT8;"`
	//[ 7] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[ 8] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
	//[ 9] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID sql.NullString `gorm:"column:delta_node_uuid;type:TEXT;"`
}

var content_miner_logsTableInfo = &TableInfo{
	Name: "content_miner_logs",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "content",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "Content",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "miner",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Miner",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "miner",
			ProtobufFieldName:  "miner",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "node_info",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "NodeInfo",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_info",
			ProtobufFieldName:  "node_info",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "requester_info",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "RequesterInfo",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "requester_info",
			ProtobufFieldName:  "requester_info",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "requesting_api_key",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "RequestingAPIKey",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "requesting_api_key",
			ProtobufFieldName:  "requesting_api_key",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "system_content_miner_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "SystemContentMinerID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "system_content_miner_id",
			ProtobufFieldName:  "system_content_miner_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "delta_node_uuid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "DeltaNodeUUID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "delta_node_uuid",
			ProtobufFieldName:  "delta_node_uuid",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ContentMinerLogs) TableName() string {
	return "content_miner_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ContentMinerLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ContentMinerLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ContentMinerLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ContentMinerLogs) TableInfo() *TableInfo {
	return content_miner_logsTableInfo
}
