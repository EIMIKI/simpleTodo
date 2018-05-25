package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ShowTodo struct {
	Todo string
	Id   int
}
type PostTodo struct {
	Todo   string `form:"new_todo"`
	Delete []int  `form:"delete[]"`
}

func getEnv() (string, string) {
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	return dbuser, dbpass
}

func openDB() *sql.DB {
	dbuser, dbpass := getEnv()
	db, err := sql.Open("mysql", ""+dbuser+":"+dbpass+"@/todos")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func Select() []ShowTodo {
	showTodos := []ShowTodo{}
	db := openDB()
	defer db.Close()
	rows, err := db.Query("select * from todo")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var showTodo ShowTodo
		err := rows.Scan(&(showTodo.Id), &(showTodo.Todo))
		if err != nil {
			log.Fatalln(err)
		}
		showTodos = append(showTodos, showTodo)
	}
	return showTodos
}

func ChangeData(postTodo PostTodo) {
	db := openDB()
	defer db.Close()
	fmt.Println(postTodo.Todo)
	//新規TODOがある場合
	if postTodo.Todo != "" {
		_, err := db.Exec("insert into todo(todo) values('" + postTodo.Todo + "');")
		if err != nil {
			log.Fatalln(err)

		}
	}
	//消去TODOがある場合
	if postTodo.Delete != nil {
		for _, v := range postTodo.Delete {
			_, err := db.Exec("delete from todo where todoid=" + strconv.Itoa(v) + ";")
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
