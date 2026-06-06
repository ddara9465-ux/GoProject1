package http

import (
	"GoProject1/internal/application/usecases/masters"
	"GoProject1/internal/infrastructure/cache"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MastersHandler struct {
	masterService *masters.MasterService
	mastersCache  *cache.MastersCache
}

// Обновляем конструктор
func NewMastersHandler(masterService *masters.MasterService, mastersCache *cache.MastersCache) *MastersHandler {
	return &MastersHandler{
		masterService: masterService,
		mastersCache:  mastersCache,
	}
}

func (h *MastersHandler) MastersList(c *gin.Context) {
	mastersData, _ := h.masterService.GetMasters(c.Request.Context())
	c.HTML(http.StatusOK, "masters.html", gin.H{"masters": mastersData})
}

func (h *MastersHandler) CreateMaster(c *gin.Context) {
	name := c.PostForm("name")
	specialization := c.PostForm("specialization")

	err := h.masterService.CreateMaster(c.Request.Context(), name, specialization)
	if err != nil {
		log.Printf("Ошибка создания мастера: %v", err)
		c.Redirect(http.StatusSeeOther, "/admin/masters")
		return
	}

	// Инвалидируем кеш после добавления мастера
	h.mastersCache.Invalidate()

	c.Redirect(http.StatusSeeOther, "/admin/masters")
}

func (h *MastersHandler) DeleteMaster(c *gin.Context) {
	masterIDStr := c.PostForm("master_id")
	masterID, err := strconv.Atoi(masterIDStr)
	if err != nil {
		log.Print("Неверный ID мастера")
		c.Redirect(http.StatusSeeOther, "/admin/masters")
		return
	}

	err = h.masterService.DeleteMaster(c.Request.Context(), masterID)
	if err != nil {
		log.Printf("Ошибка удаления мастера: %v", err)
		c.Redirect(http.StatusSeeOther, "/admin/masters")
		return
	}

	// Инвалидируем кеш после удаления мастера
	h.mastersCache.Invalidate()

	c.Redirect(http.StatusSeeOther, "/admin/masters")
}
