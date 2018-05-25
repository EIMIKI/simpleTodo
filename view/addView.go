package view

import (
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var newTodo model.NewTodo
	if c.ShouldBind(&newTodo) == nil {
		model.AddTodo(newTodo)
	}
	GetRoot(c)
}
