package model

import (
	"fmt"
	"log"
)

type NewTodo struct { //追加用
	Todo string `form:"new_todo"`
}

func AddTodo(newTodo NewTodo) {
	db := openDB()
	defer db.Close()
	fmt.Println(newTodo.Todo)
	if newTodo.Todo != "" {
		_, err := db.Exec("insert into todo(todo) values('" + newTodo.Todo + "');")
		if err != nil {
			log.Fatalln(err)

		}
	}
}
