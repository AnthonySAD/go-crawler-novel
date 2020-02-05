package shubao69

import (
	"crawler/db"
	"reflect"
)

func HandleNovel(content interface{}){
	if content == nil{
		return
	}
	DB := db.Db()
	val := reflect.ValueOf(content)
	length := val.Len()
	for i := 0; i < length; i++ {
		DB.Create(val.Index(i).Interface())
	}

}
