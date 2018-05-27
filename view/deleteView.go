package view

import (
	"net/http"
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
)

func DeleteTodo(c *gin.Context) {
	cookie := CookieCheck(c)
	var delTodo model.DelTodo
	if c.ShouldBind(&delTodo) == nil {
		model.DeleteTodo(delTodo, objx.MustFromBase64(cookie))
	}
	c.Redirect(http.StatusMovedPermanently, "/list")

}
