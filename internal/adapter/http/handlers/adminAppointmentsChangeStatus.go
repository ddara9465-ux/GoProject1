package http

import (
	"GoProject1/internal/application/usecases/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminAppointmentsHandler struct {
	adminService *admin.AdminService
}

func NewAdminAppointmentsHandler(adminService *admin.AdminService) *AdminAppointmentsHandler {
	return &AdminAppointmentsHandler{adminService: adminService}
}

func (h *AdminAppointmentsHandler) AdminUpdateAppointmentStatus(c *gin.Context) {
	idStr := c.PostForm("id")
	status := c.PostForm("status")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	if err := h.adminService.UpdateAppointmentStatus(c.Request.Context(), id, status); err != nil {
		return
	}

	c.Redirect(302, "/admin")
}
