package models

import (
	"github.com/jinzhu/gorm"

	"log"
)

var DB *gorm.DB

func InitDb() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/jwt_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err.Error())
		panic("Failed to connect to Database")
	}

	database.AutoMigrate(&Person{})
	DB = database
}
