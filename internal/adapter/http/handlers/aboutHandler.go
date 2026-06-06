package http

import (
	"GoProject1/internal/application/usecases/masters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AboutHandler struct {
	masterService *masters.MasterService
}

func NewAboutHandler(masterService *masters.MasterService) *AboutHandler {
	return &AboutHandler{masterService: masterService}
}

func (h *AboutHandler) AboutGET(c *gin.Context) {
	_, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	mastersData, _ := h.masterService.GetMasters(c.Request.Context())
	c.HTML(http.StatusOK, "about.html", gin.H{
		"masters": mastersData,
	})
}
