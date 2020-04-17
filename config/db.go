package config

import (
	"github.com/alvinarthas/learn_gin/models"
	"github.com/jinzhu/gorm"

	// set mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB is for initialization Connection
var DB *gorm.DB

// InitDB is
func InitDB() {
	var err error

	// Setting Database MYSQL
	DB, err = gorm.Open("mysql", "root:@/learn_gin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")

	}

	// Migrate the Database
	DB.AutoMigrate(&models.Article{})
}
