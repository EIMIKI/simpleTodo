package main

import (
	"os"
	"simpleTodo/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	r := controller.Init()
	gomniauth.SetSecurityKey(os.Getenv("SECUKEY")) //任意の文字列
	gomniauth.WithProviders(                       //認証に使う情報
		google.New(os.Getenv("CLIKEY"), os.Getenv("SECKEY"), "http://localhost:8080/callback"),
	)
	r.Run()
}
