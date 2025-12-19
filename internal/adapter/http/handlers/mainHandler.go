package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainGET(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", nil)
}
