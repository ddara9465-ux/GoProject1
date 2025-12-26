package http

import (
	"GoProject1/internal/adapter/persistence"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminLoginGET показывает страницу входа для админов
func AdminLoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}

// AdminLoginPOST обрабатывает вход админа
func AdminLoginPOST(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Проверяем логин/пароль
	success, adminID, err := persistence.CheckAdminAuth(username, password)
	if err != nil {
		log.Printf("Ошибка: %v", err)
	}
	if success {
		// Устанавливаем специальную куку для админа
		c.SetCookie("admin_id", strconv.Itoa(adminID), 3600, "/admin", "", false, true)
		c.SetCookie("admin_name", username, 3600, "/admin", "", false, true)

		log.Printf("Админ %s вошел в систему", username)
		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}

	// Если ошибка - показываем форму снова
	c.HTML(http.StatusOK, "admin_login.html", gin.H{
		"Error": "Неверное имя администратора или пароль",
	})
}

// AdminLogout выходит из админки
func AdminLogout(c *gin.Context) {
	// Удаляем куки админа
	c.SetCookie("admin_id", "", -1, "/admin", "", false, true)
	c.SetCookie("admin_name", "", -1, "/admin", "", false, true)

	c.Redirect(http.StatusSeeOther, "/")
}
