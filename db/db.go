package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var maxOpenConnects int = 3
var maxIdleConnects int = 1
var db *gorm.DB

func Db() *gorm.DB{
	if db == nil {
		DB, err := gorm.Open("mysql", "root:123456@/novel?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("failed to connect database")
		}
		db = DB
		db.DB().SetMaxOpenConns(maxOpenConnects)
		db.DB().SetMaxIdleConns(maxIdleConnects)
	}

	return db
}
