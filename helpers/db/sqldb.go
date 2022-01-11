package DB

import (
	"gorm.io/driver/sqlite"
	"log"
	"path/filepath"
	"thh/helpers"
	"thh/helpers/config"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	// GORM 的 MSYQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var sqlDBIns *gorm.DB

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
		sqlDBIns, err = connectSqlLiteDB()
		break
	case "mysql":
		sqlDBIns, err = connectMysqlDB()
		break
	default:
		sqlDBIns, err = connectSqlLiteDB()
		break
	}

	if err != nil {
		log.Println(err)
	}

	sqlDB, _ := sqlDBIns.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	return sqlDBIns
}

func SqlDBIns() *gorm.DB {
	return sqlDBIns
}

func connectMysqlDB() (*gorm.DB, error) {
	// 初始化 MySQL 连接信息
	var (
		dbUrl = config.GetString("database.mysql.url")
		debug = config.GetBool("app.debug")
	)

	gormConfig := mysql.New(mysql.Config{
		DSN: dbUrl,
	})

	var level gormlogger.LogLevel

	if debug {
		// 读取不到数据也会显示
		level = gormlogger.Info
	} else {
		// 只有错误才会显示
		level = gormlogger.Error
	}

	// 准备数据库连接池
	db, err := gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})
	if debug && err == nil {
		db = db.Debug()
	}
	return db, err
}
func connectSqlLiteDB() (*gorm.DB, error) {
	dbPath := config.GetString("database.sqlite.path")
	dbDir := filepath.Dir(dbPath)

	if !helpers.IsExist(dbPath) {
		if err := helpers.MkdirAll(dbDir, 0777); err != nil {

		}
		if err := helpers.FilePutContents(dbPath, []byte(""), false); err != nil {

		}
	}
	// ":memory:"
	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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
