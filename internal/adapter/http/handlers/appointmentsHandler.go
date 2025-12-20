package http

import (
	"GoProject1/internal/application/usecases/appointments"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRequestAppointments — создаёт запись на услугу (данные берём из формы + user_id из cookie).
func CreateRequestAppointments(c *gin.Context) {
	// Достаём user_id из cookie, без него не понимаем, какой клиент создаёт запись.
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		log.Print("Нет кука user_id")
		return
	}

	// user_id в cookie строкой, переводим в int для usecase/БД.
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return
	}

	// Читаем поля формы
	date := c.PostForm("date")
	employee := c.PostForm("master")
	procedure := c.PostForm("service")
	notes := c.PostForm("comment")

	// Передаём данные в слой usecases
	appointments.UC_CreateRequestAppointments(date, employee, procedure, notes, userID)
}
