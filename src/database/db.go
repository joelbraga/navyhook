package database

import (
	_ "github.com/andrepinto/navyhook/_vendor/src/github.com/mattn/go-sqlite3"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func init(){
	DB = GetDB()
	DB.AutoMigrate(&Action{})
	DB.AutoMigrate(&Configuration{})
	DB.AutoMigrate(&Repository{})
	DB.AutoMigrate(&Hook{})
}


func initDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal("Failed to open database!")
	}
	db.LogMode(true)

	return &db
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = initDb()
	}

	return DB
}