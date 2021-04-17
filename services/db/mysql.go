package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // for MySQL instances
	_ "github.com/jinzhu/gorm/dialects/sqlite" // for SQLite instances
)

// MySQL global database instance
var MySQL *gorm.DB

// ConnectMySQL init connection to MySQL database
func ConnectMySQL(databaseType, dsn string) {
	database, err := gorm.Open(databaseType, dsn)
	if err != nil {
		panic(err)
	}
	MySQL = database
}

// AutoMigration migrates a model to database
func AutoMigration(model interface{}) {
	MySQL.AutoMigrate(model)
}
