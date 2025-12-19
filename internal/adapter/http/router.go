package http

import (
	http "GoProject1/internal/adapter/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/login", http.LoginGET)
	r.POST("/login", http.LoginPOST)

	r.GET("/register", http.RegisterGET)
	r.POST("/register", http.RegisterPOST)

	r.GET("/", http.MainGET)
}
