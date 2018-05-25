package view

import (
	"net/http"
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) { //Getに対するレスポンス兼表示
	showTodos := model.Select()
	c.HTML(http.StatusOK, "index.html", showTodos)
}
