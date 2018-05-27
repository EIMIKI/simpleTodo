package model

import (
	"log"
	"strconv"

	"github.com/stretchr/objx"
)

type DelTodo struct { //削除用
	Delete []int `form:"delete[]"`
}

func DeleteTodo(delTodo DelTodo, cookie objx.Map) {
	db := openDB()
	defer db.Close()
	name := cookie["Name"].(string)

	if delTodo.Delete != nil {
		for _, v := range delTodo.Delete {
			_, err := db.Exec("delete from todo where todoid=" + strconv.Itoa(v) + "and user='" + name + "';")
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
