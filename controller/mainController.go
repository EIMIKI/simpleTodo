package controller

import "github.com/gin-gonic/gin"
import "simpleTodo/view"

func Init() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", view.GetRoot)
	r.POST("/", view.PostRoot)
	return r
}
