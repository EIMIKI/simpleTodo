package main

import (
	"simpleTodo/controller"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := controller.Init()
	r.Run()
}
