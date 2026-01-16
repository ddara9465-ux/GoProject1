package http

import (
	"GoProject1/internal/application/usecases/masters"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MastersList(c *gin.Context) {
	c.HTML(http.StatusOK, "masters.html", gin.H{"masters": masters.UC_GetMasters()})
}

func CreateMaster(c *gin.Context) {
	name := c.PostForm("name")
	specialization := c.PostForm("specialization")
	masters.UC_CreateMaster(name, specialization)
	c.Redirect(http.StatusSeeOther, "/admin/masters")
}

func DeleteMaster(c *gin.Context) {
	masterIDStr := c.PostForm("master_id")
	masterID, err := strconv.Atoi(masterIDStr)
	if err != nil {
		log.Print("Неверный ID мастера")
		return
	}
	masters.UC_DeleteMaster(masterID)
	c.Redirect(http.StatusSeeOther, "/admin/masters")
}
