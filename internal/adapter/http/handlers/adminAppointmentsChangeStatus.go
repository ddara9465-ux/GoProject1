package http

import (
	"GoProject1/internal/adapter/persistence"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Меняет статус записи по id (данные приходят из формы админки).
func AdminUpdateAppointmentStatus(c *gin.Context) {
	idStr := c.PostForm("id")
	status := c.PostForm("status")

	// id из формы -> int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	// UPDATE в БД
	if err := persistence.A_UpdateAppointmentStatus(c.Request.Context(), id, status); err != nil {
		return
	}

	// Назад в админку
	c.Redirect(302, "/admin")
}
