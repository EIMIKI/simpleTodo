package model

import (
	"log"
	"unicode/utf8"

	"github.com/stretchr/objx"
)

type NewTodo struct { //追加用
	Todo string `form:"new_todo"`
}

func AddTodo(newTodo NewTodo, cookie objx.Map) {
	db := openDB()
	defer db.Close()
	name := cookie["Name"].(string)
	if newTodo.Todo != "" {
		if utf8.RuneCountInString(newTodo.Todo) > 255 {
			newTodo.Todo = "error! : max number of charactors is 255"
		}
		_, err := db.Exec("insert into todo(todo) values('" + newTodo.Todo + "','" + name + "');")
		if err != nil {
			log.Fatalln(err)

		}
	}
}
