package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ShowTodo struct { //表示用
	Todo string
	Id   int
}
type PostTodo struct { //追加・削除用
	Todo   string `form:"new_todo"`
	Delete []int  `form:"delete[]"`
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

func Select() []ShowTodo { //表示用データの作成
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

func ChangeData(postTodo PostTodo) { //データの削除・追加
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
