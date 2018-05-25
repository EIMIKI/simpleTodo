package model

import (
	"fmt"
	"log"
	"unicode/utf8"
)

type NewTodo struct { //追加用
	Todo string `form:"new_todo"`
}

func AddTodo(newTodo NewTodo) {
	db := openDB()
	defer db.Close()
	fmt.Println(newTodo.Todo)
	if newTodo.Todo != "" {
		if utf8.RuneCountInString(newTodo.Todo) > 255 {
			newTodo.Todo = "error! : max number of charactors is 255"
		}
		_, err := db.Exec("insert into todo(todo) values('" + newTodo.Todo + "');")
		if err != nil {
			log.Fatalln(err)

		}
	}
}
