package http

import (
	"GoProject1/internal/application/usecases/appointments"

	"github.com/gin-gonic/gin"
)

func CreateRequestAppointments(c *gin.Context) {
	date := c.PostForm("date")         // Желаемая дата
	employee := c.PostForm("master")   // Мастер
	procedure := c.PostForm("service") // Услуга
	notes := c.PostForm("comment")
	appointments.UC_CreateRequestAppointments(date, employee, procedure, notes)
}
