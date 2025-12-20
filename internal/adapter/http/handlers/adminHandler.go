package http

import (
	"GoProject1/internal/application/usecases/admin"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	//проверяем кук входа, если нет - перенаправляет на /login
	//читаем куки
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		// куки нет
		log.Print("Нет кука входа")
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// переводим кук в число(была строка)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return
	}

	if admin.UC_Accesadmin(userID) {
		log.Print("Открытие admin.html")
		apps := admin.UC_getAppointments() // получаем слайс из БД
		c.HTML(http.StatusOK, "admin.tmpl", gin.H{
			"Appointments": apps, // сюда слайс из БД
		})
	} else {
		c.Redirect(http.StatusSeeOther, "/") // если не админ - перенаправляем на главную страницу
	}
}
