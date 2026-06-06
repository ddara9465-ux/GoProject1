package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *AdminAppointmentsHandler) AdminEditDetails(c *gin.Context) {
	idStr := c.PostForm("id")
	date := c.PostForm("date")
	master := c.PostForm("master")
	service := c.PostForm("service")

	id, err := strconv.Atoi(idStr)
	if err != nil || date == "" || master == "" || service == "" {
		log.Printf("Неверные данные: id=%s, date=%s, master=%s, service=%s",
			idStr, date, master, service)
		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}

	err = h.adminService.UpdateAppointmentDetails(c.Request.Context(), id, date, master, service)
	if err != nil {
		log.Printf("Ошибка обновления записи %d: %v", id, err)
	}

	c.Redirect(http.StatusSeeOther, "/admin")
}
