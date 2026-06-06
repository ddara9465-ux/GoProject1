package http

import (
	"GoProject1/internal/application/usecases/admin"
	"GoProject1/internal/application/usecases/masters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService  *admin.AdminService
	masterService *masters.MasterService
}

func NewAdminHandler(adminService *admin.AdminService, masterService *masters.MasterService) *AdminHandler {
	return &AdminHandler{
		adminService:  adminService,
		masterService: masterService,
	}
}

func (h *AdminHandler) AdminMain(c *gin.Context) {
	_, err := c.Cookie("isAdmin")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	ctx := c.Request.Context()
	apps := h.adminService.GetAppointments(ctx)
	mastersData, _ := h.masterService.GetMasters(ctx)

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"Appointments": apps,
		"Masters":      mastersData,
	})
}
