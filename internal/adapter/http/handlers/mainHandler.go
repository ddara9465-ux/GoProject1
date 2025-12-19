package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainGET(c *gin.Context) {
	//проверяем кук входа, если нет - перенаправляет на /login
	_, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	} else {
		c.HTML(http.StatusOK, "main.html", nil)
	}
}
