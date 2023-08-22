package database

import (
	"belajar-fiber-gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/go?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{})

	DB = db
}
