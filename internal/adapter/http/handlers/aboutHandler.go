package http

import (
	"GoProject1/internal/application/usecases/masters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutGET(c *gin.Context) {
	_, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	mastersData := masters.UC_GetMasters()

	c.HTML(http.StatusOK, "about.html", gin.H{
		"masters": mastersData,
	})
}
