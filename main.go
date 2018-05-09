package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//POST受け用
type PostTodo struct {
	Todo   string `form:"new_todo"`
	Delete []int  `form:"delete[]"`
}

//リスト表示用
type ShowTodo struct {
	Todo string
	Id   int
}

//現在のTodoリストの表示
func getRoot(c *gin.Context) {
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
	c.HTML(http.StatusOK, "index.html", showTodos)
}

// POSTに関する処理
func postRoot(c *gin.Context) {
	var postTodo PostTodo
	if c.ShouldBind(&postTodo) == nil {
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
	getRoot(c)
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "yourUsername:yourPassword@/todos")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", getRoot)
	r.POST("/", postRoot)
	r.Run()
}
