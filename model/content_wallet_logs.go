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


Table: content_wallet_logs
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 2] wallet                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] system_content_wallet_id                       INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 7] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 8] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 9] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] wallet_id                                      INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 32,    "content": 30,    "wallet": "IPWZJQDFqOtrZAiOqxXsTcxGA",    "nodeInfo": "bqfDOZUYMqlEMJMpMNWWCCAFU",    "requesterInfo": "RGnuoAwqrHeiSIbCNaAtinoEn",    "requestingApiKey": "oKedRqCcjjdtkhmRKWWoKYqkN",    "systemContentWalletId": 40,    "createdAt": "2253-07-27T03:41:17.410938197-04:00",    "updatedAt": "2272-09-27T09:42:15.555605978-04:00",    "deltaNodeUuid": "MmyGprUgsHYiZKBoEcxYVUCrN",    "walletId": 32}



*/

// ContentWalletLogs struct is a row record of the content_wallet_logs table in the estuary database
type ContentWalletLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Content null.Int `gorm:"column:content;type:INT8;"`
	//[ 2] wallet                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Wallet null.String `gorm:"column:wallet;type:TEXT;"`
	//[ 3] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo null.String `gorm:"column:node_info;type:TEXT;"`
	//[ 4] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo null.String `gorm:"column:requester_info;type:TEXT;"`
	//[ 5] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequestingAPIKey null.String `gorm:"column:requesting_api_key;type:TEXT;"`
	//[ 6] system_content_wallet_id                       INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemContentWalletID null.Int `gorm:"column:system_content_wallet_id;type:INT8;"`
	//[ 7] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[ 8] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
	//[ 9] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID null.String `gorm:"column:delta_node_uuid;type:TEXT;"`
	//[10] wallet_id                                      INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	WalletID null.Int `gorm:"column:wallet_id;type:INT8;"`
}

var content_wallet_logsTableInfo = &TableInfo{
	Name: "content_wallet_logs",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "wallet",
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
			GoFieldName:        "Wallet",
			GoFieldType:        "null.String",
			JSONFieldName:      "wallet",
			ProtobufFieldName:  "wallet",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "nodeInfo",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "requesterInfo",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "requestingApiKey",
			ProtobufFieldName:  "requesting_api_key",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "system_content_wallet_id",
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
			GoFieldName:        "SystemContentWalletID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "systemContentWalletId",
			ProtobufFieldName:  "system_content_wallet_id",
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
			GoFieldType:        "null.Time",
			JSONFieldName:      "createdAt",
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
			GoFieldType:        "null.Time",
			JSONFieldName:      "updatedAt",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "deltaNodeUuid",
			ProtobufFieldName:  "delta_node_uuid",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "wallet_id",
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
			GoFieldName:        "WalletID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "walletId",
			ProtobufFieldName:  "wallet_id",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ContentWalletLogs) TableName() string {
	return "content_wallet_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ContentWalletLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ContentWalletLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ContentWalletLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ContentWalletLogs) TableInfo() *TableInfo {
	return content_wallet_logsTableInfo
}
