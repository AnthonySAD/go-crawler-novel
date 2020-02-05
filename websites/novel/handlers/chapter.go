package shubao69

import (
	"crawler/db"
	"reflect"
)

func HandleChapter(content interface{}){

	if reflect.ValueOf(content).IsNil() {
		return
	}

	DB := db.Db()
	DB.Create(content)
}
