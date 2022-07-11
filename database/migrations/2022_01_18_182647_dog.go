package migrations

import (
	"database/sql"
	"thh/arms/migrate"
	"time"

	"gorm.io/gorm"
)

func init() {
	type Tmp struct {
		ID        uint `gorm:"primaryKey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
		Name1     string         `gorm:"type:varchar(255);not null;default:'';"  json:"username"`
		Name2     string         `gorm:"type:varchar(255);index:idx_email,unique;default:'';"  json:"email"`
	}
	migrate.Add("2022_01_18_182647_dog",
		func(migrator gorm.Migrator, DB *sql.DB) {
			migrator.AutoMigrate(&Tmp{})
		},
		func(migrator gorm.Migrator, DB *sql.DB) {
			migrator.DropTable(&Tmp{})
		})
}
