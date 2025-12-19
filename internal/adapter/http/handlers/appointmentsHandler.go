package http

import (
	"GoProject1/internal/application/usecases/appointments"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRequestAppointments(c *gin.Context) {
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		// куки нет
		log.Print("Нет кука входа")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id cookie"})
		return
	}

	date := c.PostForm("date")         // Желаемая дата
	employee := c.PostForm("master")   // Мастер
	procedure := c.PostForm("service") // Услуга
	notes := c.PostForm("comment")
	appointments.UC_CreateRequestAppointments(date, employee, procedure, notes, userID)
}
