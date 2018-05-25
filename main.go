package main

import (
	"database/sql"
	"log"
	"os"
	"simpleTodo/controller"

	_ "github.com/go-sql-driver/mysql"
)

func openDB() *sql.DB {
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	db, err := sql.Open("mysql", ""+dbuser+":"+dbpass+"@/todos")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	r := controller.Init()
	r.Run()
}
