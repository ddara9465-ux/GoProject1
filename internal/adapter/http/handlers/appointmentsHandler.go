package http

import (
	"GoProject1/internal/application/usecases/appointments"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRequestAppointments(c *gin.Context) {
	//читаем куки
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		// куки нет
		log.Print("Нет кука входа")
		return
	}
	// переводим кук в число(была строка)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return
	}

	date := c.PostForm("date")         // Желаемая дата
	employee := c.PostForm("master")   // Мастер
	procedure := c.PostForm("service") // Услуга
	notes := c.PostForm("comment")
	appointments.UC_CreateRequestAppointments(date, employee, procedure, notes, userID)
}
