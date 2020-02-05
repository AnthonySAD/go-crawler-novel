package db

import shubao69 "crawler/websites/shubao69/models"

func Init(){
	DB := Db()
	DB.AutoMigrate(&shubao69.Novel{})
	DB.AutoMigrate(&shubao69.Chapter{})
	DB.Close()
}
