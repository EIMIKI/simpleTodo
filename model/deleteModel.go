package model

import (
	"log"
	"strconv"
)

type DelTodo struct { //削除用
	Delete []int `form:"delete[]"`
}

func DeleteTodo(delTodo DelTodo) {
	db := openDB()
	defer db.Close()

	if delTodo.Delete != nil {
		for _, v := range delTodo.Delete {
			_, err := db.Exec("delete from todo where todoid=" + strconv.Itoa(v) + ";")
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
