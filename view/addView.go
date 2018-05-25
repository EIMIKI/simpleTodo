package view

import (
	"net/http"
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var newTodo model.NewTodo
	if c.ShouldBind(&newTodo) == nil {
		model.AddTodo(newTodo)
	}
	c.Redirect(http.StatusMovedPermanently, "/list")
}
