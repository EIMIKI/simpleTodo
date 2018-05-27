package view

import (
	"log"
	"net/http"
	"simpleTodo/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
)

func ShowList(c *gin.Context) { //Getに対するレスポンス兼表示
	if cookie, err := c.Cookie("simpletodo"); err == http.ErrNoCookie {
		c.Redirect(http.StatusMovedPermanently, "/login")
	} else if err != nil {
		log.Fatalln(err)
	} else {
		showTodos := model.Select(objx.MustFromBase64(cookie))
		c.HTML(http.StatusOK, "index.html", showTodos)
	}
}
