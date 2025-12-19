package http

import (
	"GoProject1/internal/application/usecases/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGET(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegisterPOST(c *gin.Context) {
	// получаеи данные с формы
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	password := c.PostForm("password")

	//функция регистрации
	if auth.Register(password, first_name, last_name, phone, email) {
		log.Print("Успешная регистрация")
	} else {
		log.Print("Ошибка регистрации")
	}
}
