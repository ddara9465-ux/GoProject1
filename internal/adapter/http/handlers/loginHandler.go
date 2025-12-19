package http

import (
	"GoProject1/internal/application/usecases/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPOST(c *gin.Context) {
	// получаем данные с формы
	login := c.PostForm("login")
	password := c.PostForm("password")

	// проверка корректности логина и пароля
	if auth.UC_Login(login, password) {
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		c.HTML(200, "login.html", gin.H{
			"Error": "Неверный логин или пароль",
		})
	}
}
