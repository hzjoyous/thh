package dbconnect

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/spf13/cast"
	"log"
	"os"
	"path/filepath"
	"thh/arms"
	"thh/arms/config"
	"thh/arms/logger"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	// GORM 的 MSYQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var dbIns *gorm.DB
var SQLDB *sql.DB

// NewMysql dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
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
		fmt.Println("use sqlite")
		dbIns, err = connectSqlLiteDB(logger.NewGormLogger())
		break
	case "mysql":
		fmt.Println("use mysql")
		dbIns, err = connectMysqlDB(logger.NewGormLogger())
		break
	default:
		fmt.Println("use sqlite")
		dbIns, err = connectSqlLiteDB(logger.NewGormLogger())
		break
	}

	if err != nil {
		log.Println(err)
		panic(err)
	}

	// 获取底层的 sqlDB
	SQLDB, _ = dbIns.DB()
	// 设置最大连接数
	SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	return dbIns
}

func Std() *gorm.DB {
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
		Logger: _logger,
	})
	if debug && err == nil {
		fmt.Println("开启debug")
		db = db.Debug()
	}
	return db, err
}
func connectSqlLiteDB(_logger gormlogger.Interface) (*gorm.DB, error) {
	dbPath := config.GetString("database.sqlite.path")
	debug := config.GetBool("app.debug")
	dbDir := filepath.Dir(dbPath)
	if dbPath == ":memory:" {

	} else if !arms.IsExist(dbPath) {
		if err := os.MkdirAll(dbDir, 0777); err != nil {

		}
		if err := arms.FilePutContents(dbPath, []byte(""), false); err != nil {

		}
	}
	// ":memory:"
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: _logger})
	if debug && err == nil {
		fmt.Println("开启debug")
		db = db.Debug()
	}
	return db, err
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
	return cast.ToString(a.ID)
}

func TableName(obj any) string {
	stmt := &gorm.Statement{DB: dbIns}
	stmt.Parse(obj)
	return stmt.Schema.Table
}

func CurrentDatabase() (dbname string) {
	dbname = Std().Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {

	var err error

	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMysqlDatabase()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteAllSqliteTables() error {
	var tables []string
	Std().Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'")
	for _, table := range tables {
		return Std().Migrator().DropTable(table)
	}
	return nil
}

func deleteMysqlDatabase() error {
	dbname := CurrentDatabase()
	sqlStr := fmt.Sprintf("DROP DATABASE %s;", dbname)
	if err := Std().Exec(sqlStr).Error; err != nil {
		return err
	}
	sqlStr = fmt.Sprintf("CREATE DATABASE %s;", dbname)
	if err := Std().Exec(sqlStr).Error; err != nil {
		return err
	}
	sqlStr = fmt.Sprintf("USE %s;", dbname)
	if err := Std().Exec(sqlStr).Error; err != nil {
		return err
	}
	return nil
}
