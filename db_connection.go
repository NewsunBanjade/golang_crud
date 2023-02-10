package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDns = "root:root@tcp(localhost:3306)/mydb?parseTime=true"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDns), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed")
	}
	Database.AutoMigrate(&Employee{})

}
