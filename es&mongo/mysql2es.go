package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const batchSize = 2000 // 每批次处理的数据量

// ColumnInfo holds information about a column
type ColumnInfo struct {
	ColumnName string `gorm:"column:COLUMN_NAME"`
	DataType   string `gorm:"column:DATA_TYPE"`
	ColumnType string `gorm:"column:COLUMN_TYPE"`
	ColumnKey  string `gorm:"column:COLUMN_KEY"`
	Extra      string `gorm:"column:EXTRA"`
}

// Map MySQL data types to Elasticsearch data types
func mapMySQLToElasticType(mysqlType string) string {
	switch mysqlType {
	case "int", "tinyint", "smallint", "mediumint", "bigint":
		return "integer"
	case "float", "double", "decimal":
		return "float"
	case "varchar", "char", "text", "mediumtext", "longtext":
		return "text"
	case "date", "datetime", "timestamp":
		return "date"
	case "blob", "binary", "varbinary":
		return "binary"
	default:
		return "text"
	}
}

type Order struct {
	ID                int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderNum          int64      `gorm:"column:order_num" json:"order_num"`
	SkuID             int64      `gorm:"column:sku_id" json:"sku_id"`
	CarType           int        `gorm:"column:car_type" json:"car_type"`
	CarCollectionType int        `gorm:"column:car_collection_type" json:"car_collection_type"`
	EvaluatorUID      int64      `gorm:"column:evaluator_uid" json:"evaluator_uid"`
	CollectionTime    time.Time  `gorm:"column:collection_time" json:"collection_time"`
	CollectionAddress string     `gorm:"column:collection_address" json:"collection_address"`
	ContactName       string     `gorm:"column:contact_name" json:"contact_name"`
	ContactPhone      string     `gorm:"column:contact_phone" json:"contact_phone"`
	SkuImportInfo     string     `gorm:"column:sku_import_info" json:"sku_import_info"`
	TradeOrderBiz     string     `gorm:"column:trade_order_biz" json:"trade_order_biz"`
	CarInfo           string     `gorm:"column:car_info" json:"car_info"`
	TransferInfo      string     `gorm:"column:transfer_info" json:"transfer_info"`
	FweTradeOrderID   string     `gorm:"column:fwe_trade_order_id" json:"fwe_trade_order_id"`
	Extra             string     `gorm:"column:extra" json:"extra"`
	CreateTime        time.Time  `gorm:"column:create_time" json:"create_time"`
	UpdatedTime       time.Time  `gorm:"column:updated_time" json:"updated_time"`
	IsTest            uint8      `gorm:"column:is_test" json:"is_test"`
	ShopID            int64      `gorm:"column:shop_id" json:"shop_id"`
	Remark            string     `gorm:"column:remark" json:"remark"`
	TradeServiceType  int        `gorm:"column:trade_service_type" json:"trade_service_type"`
	AuctionInfo       string     `gorm:"column:auction_info" json:"auction_info"`
	OrderSource       int        `gorm:"column:order_source" json:"order_source"`
	FsmOrderStatus    int        `gorm:"column:fsm_order_status" json:"fsm_order_status"`
	AuctionCount      *int       `gorm:"column:auction_count" json:"auction_count"`
	CollectCount      *int       `gorm:"column:collect_count" json:"collect_count"`
	FinishTime        *time.Time `gorm:"column:finish_time" json:"finish_time"`
	LockVersion       *int       `gorm:"column:lock_version" json:"lock_version"`
	DefeatedInfo      string     `gorm:"column:defeated_info" json:"defeated_info"`
	OrderStatus       int        `gorm:"column:order_status" json:"order_status"`
}

func mysql2ES() {
	// MySQL database connection
	dsn := "user:pw@tcp(11:11:11:11:3306)/test?parseTime=true"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 详细级别
			IgnoreRecordNotFoundError: true,        // 忽略记录未找到错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)
	}

	// Get table name
	tableName := "car_collection_order" // replace with your table name

	// Get column information for the table
	var columns []ColumnInfo
	db.Raw("SELECT COLUMN_NAME, DATA_TYPE, COLUMN_TYPE, COLUMN_KEY, EXTRA FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?", "test", tableName).Scan(&columns)

	// Elasticsearch client
	esClient, err := elastic.NewClient(
		elastic.SetURL("http://11:11:11:11:9200"),
		elastic.SetSniff(false),                                               // Disable sniffing
		elastic.SetHealthcheckInterval(10*time.Second),                        // Set health check interval
		elastic.SetErrorLog(log.New(log.Writer(), "ELASTIC ", log.LstdFlags)), // Set error logger
	)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	// Construct mapping dynamically from MySQL columns
	properties := make(map[string]interface{})
	for _, column := range columns {
		properties[column.ColumnName] = map[string]string{
			"type": mapMySQLToElasticType(column.DataType),
		}
	}

	mapping := map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": properties,
		},
	}

	mappingJSON, err := json.Marshal(mapping)
	if err != nil {
		log.Fatalf("Error marshalling mapping to JSON: %v", err)
	}

	// Create Elasticsearch index with dynamic mapping
	indexName := tableName
	exists, err := esClient.IndexExists(indexName).Do(context.Background())
	if err != nil {
		log.Fatalf("Error checking if index exists: %v", err)
	}
	if !exists {
		_, err := esClient.CreateIndex(indexName).BodyString(string(mappingJSON)).Do(context.Background())
		if err != nil {
			log.Fatalf("Error creating index with mapping: %v", err)
		}
	}

	offset := 0
	// 批量写入数据到 Elasticsearch
	for {
		var results []*Order
		err = db.Table(tableName).Offset(offset).Limit(batchSize).Find(&results).Error
		if err != nil {
			log.Fatalf("Error fetching data from MySQL: %v", err)
		}

		if len(results) == 0 {
			log.Println("No more data found in the table.")
			break // 退出循环
		}

		log.Printf("Fetched %d rows from MySQL, processing batch starting at offset %d.\n", len(results), offset)

		// 创建批量处理器
		bulkProcessor, err := esClient.BulkProcessor().
			Name("bulk_processor").
			Workers(2).
			BulkActions(batchSize).
			After(func(executionId int64, requests []elastic.BulkableRequest, response *elastic.BulkResponse, err error) {
				if err != nil || response.Errors {
					log.Printf("Failed to execute bulk: %s", err)
				} else {
					log.Printf("Successfully indexed %d documents", len(requests))
				}
			}).
			Do(context.Background())
		if err != nil {
			log.Fatalf("Failed to create bulk processor: %s", err)
		}

		// 向批量处理器中添加每个文档
		for _, result := range results {
			jsonData, err := json.Marshal(result)
			if err != nil {
				log.Printf("Error marshalling row to JSON: %v", err)
				continue
			}

			req := elastic.NewBulkIndexRequest().
				Index(indexName).
				Id(strconv.FormatInt(result.ID, 10)).
				Doc(string(jsonData))

			bulkProcessor.Add(req)
		}

		// 关闭批量处理器
		if err := bulkProcessor.Close(); err != nil {
			log.Printf("Failed to close bulk processor: %s", err)
		}

		offset += batchSize
	}

	log.Println("Data synchronization from MySQL to Elasticsearch completed successfully.")
}
