package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/stretchr/objx"
)

type ShowTodo struct { //表示用
	User objx.Map
	Todo string
	Id   int
}

func getEnv() (string, string) { //環境変数の取得
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	return dbuser, dbpass
}

func openDB() *sql.DB { //データベースの接続
	dbuser, dbpass := getEnv()
	db, err := sql.Open("mysql", ""+dbuser+":"+dbpass+"@/todos")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func Select(cookie objx.Map) []ShowTodo { //表示用データの作成
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
		fmt.Println(cookie["Name"])
		showTodos = append(showTodos, showTodo)
	}
	return showTodos
}
