package http

import (
	"GoProject1/internal/application/usecases/appointments"
	"GoProject1/internal/infrastructure/telegram"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentsHandler struct {
	appointmentService *appointments.AppointmentService
}

func NewAppointmentsHandler(appointmentService *appointments.AppointmentService) *AppointmentsHandler {
	return &AppointmentsHandler{appointmentService: appointmentService}
}

func (h *AppointmentsHandler) CreateRequestAppointments(c *gin.Context) {
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		log.Print("Нет кука user_id")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return
	}

	date := c.PostForm("date")
	employee := c.PostForm("master")
	procedure := c.PostForm("service")
	notes := c.PostForm("comment")

	ctx := c.Request.Context()

	h.appointmentService.CreateRequestAppointments(ctx, date, employee, procedure, notes, userID)

	go telegram.SendNewAppointmentNotify(userID, date, employee, procedure)

	c.Redirect(http.StatusSeeOther, "/")
}
