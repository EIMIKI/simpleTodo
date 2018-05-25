package view

import (
	"net/http"
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {
	showTodos := model.Select()
	c.HTML(http.StatusOK, "index.html", showTodos)
}

func PostRoot(c *gin.Context) {
	var postTodo model.PostTodo
	if c.ShouldBind(&postTodo) == nil {
		model.ChangeData(postTodo)
	}
	GetRoot(c)
}
