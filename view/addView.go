package view

import (
	"net/http"
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
)

func AddTodo(c *gin.Context) {
	cookie := CookieCheck(c)
	if cookie != "" {
		var newTodo model.NewTodo
		if c.ShouldBind(&newTodo) == nil {
			model.AddTodo(newTodo, objx.MustFromBase64(cookie))
		}
		c.Redirect(http.StatusMovedPermanently, "/list")
	}
}
