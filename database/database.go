package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Connect() {
	dsn := "fiberTester:1234@tcp(127.0.0.1:3306)/fiber_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
}

func Get() *gorm.DB {
	return db
}
