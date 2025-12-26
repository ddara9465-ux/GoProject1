package http

import (
	"GoProject1/internal/application/usecases/auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// показать страницу логина
func LoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// отправить запрос аутентификации
func LoginPOST(c *gin.Context) {
	// получаем данные с формы
	login := c.PostForm("login")
	password := c.PostForm("password")

	// Функция проверка корректности логина и пароля
	userID, err := auth.UC_Login(login, password)
	if err {
		c.Redirect(http.StatusSeeOther, "/")
		c.SetCookie("user_id", strconv.Itoa(userID), 3600, "/", "", false, true) // устанавливаем кук с значение userID (запомнили пользователя)
	} else {
		c.HTML(200, "login.html", gin.H{
			"Error": "Неверный логин или пароль",
		})

	}
}
