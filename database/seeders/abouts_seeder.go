package seeders

import (
	"fmt"
	"thh/database/factories"
	"thh/helpers/console"
	"thh/helpers/logger"
	"thh/helpers/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAboutsTable", func(db *gorm.DB) {

		abouts  := factories.MakeAbouts(10)

		result := db.Table("abouts").Create(&abouts)

		if err := result.Error; err != nil {
			logger.Std().Warn(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}