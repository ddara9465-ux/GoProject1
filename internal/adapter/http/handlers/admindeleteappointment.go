package http

import (
	"GoProject1/internal/application/usecases/admin"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AppointmentDelete(c *gin.Context) {
	appointmentIDStr := c.PostForm("id")
	appointmentID, err := strconv.Atoi(appointmentIDStr)
	if err != nil {
		log.Printf("Неверный ID записи: %s", appointmentIDStr)
		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}
	err = admin.UC_DeleteAppointment(appointmentID)
	if err != nil {
		log.Printf("Ошибка удаления записи: %v", err)
	}
	c.Redirect(http.StatusSeeOther, "/admin")
}
