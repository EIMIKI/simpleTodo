package controller

import "github.com/gin-gonic/gin"
import "simpleTodo/view"

func Init() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/list", view.GetRoot)
	r.POST("/add", view.AddTodo)
	r.POST("/delete", view.DeleteTodo)
	return r
}
