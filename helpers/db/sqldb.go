package db

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"log"
	"path/filepath"
	"thh/helpers"
	"thh/helpers/config"
	"thh/helpers/logger"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	// GORM 的 MSYQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var dbIns *gorm.DB
var SQLDB *sql.DB

//  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// NewMysql
func NewMysql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

// ConnectDB 初始化模型
func connectDB() *gorm.DB {
	var err error
	switch config.GetString("database.default") {
	case "sqlite":
		dbIns, err = connectSqlLiteDB( logger.NewGormLogger())
		break
	case "mysql":
		dbIns, err = connectMysqlDB( logger.NewGormLogger())
		break
	default:
		dbIns, err = connectSqlLiteDB( logger.NewGormLogger())
		break
	}

	if err != nil {
		log.Println(err)
	}

	SQLDB, _ = dbIns.DB()
	// 获取底层的 sqlDB

	// 设置最大连接数
	SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	return dbIns
}

func SqlDBIns() *gorm.DB {
	return dbIns
}

func connectMysqlDB(_logger gormlogger.Interface) (*gorm.DB, error) {
	// 初始化 MySQL 连接信息
	var (
		dbUrl = config.GetString("database.mysql.url")
		debug = config.GetBool("app.debug")
	)

	gormConfig := mysql.New(mysql.Config{
		DSN: dbUrl,
	})

	// 准备数据库连接池
	db, err := gorm.Open(gormConfig, &gorm.Config{
		Logger:_logger,
	})
	if debug && err == nil {
		db = db.Debug()
	}
	return db, err
}
func connectSqlLiteDB(_logger gormlogger.Interface) (*gorm.DB, error) {
	dbPath := config.GetString("database.sqlite.path")
	dbDir := filepath.Dir(dbPath)

	if !helpers.IsExist(dbPath) {
		if err := helpers.MkdirAll(dbDir, 0777); err != nil {

		}
		if err := helpers.FilePutContents(dbPath, []byte(""), false); err != nil {

		}
	}
	// ":memory:"
	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: _logger})
}

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`

	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`

	// 支持 gorm 软删除
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return helpers.ToString(a.ID)
}

func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: dbIns}
	stmt.Parse(obj)
	return stmt.Schema.Table
}


func CurrentDatabase() (dbname string) {
	dbname = SqlDBIns().Migrator().CurrentDatabase()
	return
}


func DeleteAllTables() error {

	var err error

	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMysqlDatabase()
	case "sqlite":
		deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteAllSqliteTables() error {
	tables := []string{}
	SqlDBIns().Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'")
	for _, table := range tables {
		SqlDBIns().Migrator().DropTable(table)
	}
	return nil
}

func deleteMysqlDatabase() error {
	dbname := CurrentDatabase()
	sql := fmt.Sprintf("DROP DATABASE %s;", dbname)
	if err := SqlDBIns().Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("CREATE DATABASE %s;", dbname)
	if err := SqlDBIns().Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("USE %s;", dbname)
	if err := SqlDBIns().Exec(sql).Error; err != nil {
		return err
	}
	return nil
}