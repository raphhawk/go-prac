package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	//d, err := gorm.Open("mysql", "user:pw/simplerest?charset=utf8&parseTime=True&loc=Local")
	dsn := "root:my-secret-pw@tcp(localhost:3306)/my_db?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Println("Connected to DB")
	if err != nil {
		log.Fatalf("Error connecting to DB: %v\n", err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
