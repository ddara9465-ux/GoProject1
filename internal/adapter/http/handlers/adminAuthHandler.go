package http

import (
	"GoProject1/internal/application/usecases/auth"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Отобразить экран админ логина
func AdminLoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}

// Запрос авторизации
func AdminLoginPOST(c *gin.Context) {
	//Получаем данные с формы
	login := c.PostForm("login")
	password := c.PostForm("password")

	//Вызываем функцию проверки авторизации
	userID, err := auth.UC_AdmimAuth(login, password)
	if err != nil {
		log.Printf("Ошибка авторизации: %v", err)
		c.Redirect(http.StatusSeeOther, "/login")
	} else {
		isAdmin := 1 // Ставим что админ, если нет ошибки
		// устанавливаем кук с значение userID (запомнили пользователя)
		c.SetCookie("user_id", strconv.Itoa(userID), 3600, "/", "", false, true)

		// устанавливаем кук с значение isAdmin (запомнили админа)
		c.SetCookie("isAdmin", strconv.Itoa(isAdmin), 3600, "/", "", false, true)
		c.Redirect(http.StatusSeeOther, "/admin")
	}

}
