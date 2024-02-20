package db

import (
	"fmt"
	"github.com/konekotech/crudAPI/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() {
	// dbはコンテナ名
	// mydevはデータベース名
	dbConn := "root:password@tcp(db:3306)/mydev?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dbConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Post{})
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
}
