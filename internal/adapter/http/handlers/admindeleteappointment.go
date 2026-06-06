package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *AdminAppointmentsHandler) AppointmentDelete(c *gin.Context) {
	appointmentIDStr := c.PostForm("id")
	appointmentID, err := strconv.Atoi(appointmentIDStr)

	if err != nil {
		log.Printf("Неверный ID записи: %s", appointmentIDStr)
		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}

	err = h.adminService.DeleteAppointment(c.Request.Context(), appointmentID)
	if err != nil {
		log.Printf("Ошибка удаления записи: %v", err)
	}

	c.Redirect(http.StatusSeeOther, "/admin")
}
