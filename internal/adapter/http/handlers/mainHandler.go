package http

import (
	"GoProject1/internal/application/usecases/masters"
	"GoProject1/internal/infrastructure/cache"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainHandler struct {
	masterService *masters.MasterService
	mastersCache  *cache.MastersCache
}

func NewMainHandler(masterService *masters.MasterService, mastersCache *cache.MastersCache) *MainHandler {
	return &MainHandler{
		masterService: masterService,
		mastersCache:  mastersCache,
	}
}

func (h *MainHandler) MainGET(c *gin.Context) {
	cookie, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	log.Print("Cookie:", cookie)

	mastersData := h.mastersCache.GetMasters()
	if mastersData == nil || h.mastersCache.IsExpired() {
		mastersData, _ = h.masterService.GetMasters(c.Request.Context())
		h.mastersCache.SetMasters(mastersData)
	}

	c.HTML(http.StatusOK, "main.html", gin.H{
		"masters": mastersData,
	})
}
