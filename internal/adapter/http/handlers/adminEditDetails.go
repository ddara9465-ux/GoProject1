package http

import (
	"GoProject1/internal/application/usecases/admin"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminEditDetails обновляет детали записи
func AdminEditDetails(c *gin.Context) {
	// Получаем данные из формы
	idStr := c.PostForm("id")
	date := c.PostForm("date")
	master := c.PostForm("master")
	service := c.PostForm("service")

	// Проверяем данные
	id, err := strconv.Atoi(idStr)
	if err != nil || date == "" || master == "" || service == "" {
		log.Printf("Неверные данные: id=%s, date=%s, master=%s, service=%s",
			idStr, date, master, service)
		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}
	//Обновляем запись
	err = admin.UC_UpdateAppointmentDetails(id, date, master, service)
	if err != nil {
		log.Printf("Ошибка обновления записи %d: %v", id, err)
	}

	//Возвращаем в админку
	c.Redirect(http.StatusSeeOther, "/admin")
}
