package config

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	cfg := mysql.Config{
		User:   "admin",
		Passwd: "12345qwerty",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "bookstore_sample",
		Params: map[string]string{"charset": "utf8", "parseTime": "True", "loc": "Local"},
	}

	// Connect to the database
	d, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
