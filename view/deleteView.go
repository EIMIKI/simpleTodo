package view

import (
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	var delTodo model.DelTodo
	if c.ShouldBind(&delTodo) == nil {
		model.DeleteTodo(delTodo)
	}
	GetRoot(c)
}
