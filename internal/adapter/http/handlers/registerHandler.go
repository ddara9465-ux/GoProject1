package http

import (
	"GoProject1/internal/application/usecases/auth"
	"log"
	"net/http"
	"strconv"

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
	userID, err := auth.UC_Register(password, first_name, last_name, phone, email)
	if err {
		log.Print("Успешная регистрация")
		c.Redirect(http.StatusSeeOther, "/")
		c.SetCookie("user_id", strconv.Itoa(userID), 3600, "/", "", false, true) // устанавливаем кук с значение userID (запомнили пользователя)
	} else {
		log.Print("Ошибка регистрации")
	}
}
