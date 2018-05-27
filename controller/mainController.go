package controller

import (
	"log"
	"net/http"
	"simpleTodo/view"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")
	r.GET("/list", view.ShowList)
	r.POST("/add", view.AddTodo)
	r.POST("/delete", view.DeleteTodo)
	r.GET("/login", Login)
	r.GET("/callback", Callback)
	return r
}

func Login(c *gin.Context) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		log.Fatalln("プロバイダの取得に失敗", err)
	}

	url, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	c.HTML(http.StatusOK, "login.html", url)
}

func Callback(c *gin.Context) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		log.Fatalln(err)
	}

	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(c.Request.URL.RawQuery))
	if err != nil {
		log.Fatalln(err)
	}

	user, err := provider.GetUser(creds)
	if err != nil {
		log.Fatalln(err)
	}

	cookieValue := objx.New(map[string]interface{}{
		"Id": user.AuthCode(),
	}).MustBase64()

	//c.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
	c.SetCookie("simpletodo", cookieValue, 10000, "/", "", false, false)
	c.Redirect(http.StatusSeeOther, "/list")
}
