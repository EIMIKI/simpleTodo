package view

import (
	"log"
	"net/http"
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
)

func ShowList(c *gin.Context) { //Getに対するレスポンス兼表示
	cookie := CookieCheck(c)
	showTodos := model.Select(objx.MustFromBase64(cookie))
	c.HTML(http.StatusOK, "index.html", showTodos)

}

func CookieCheck(c *gin.Context) string {
	cookie, err := c.Cookie("simpletodo")
	if err == http.ErrNoCookie {
		c.Redirect(http.StatusSeeOther, "/login")
	} else if err != nil {
		log.Fatalln(err)
	}

	return cookie

}
