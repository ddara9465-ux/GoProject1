package http

import (
	"GoProject1/internal/application/usecases/auth"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminAuthHandler struct {
	authService *auth.AuthService
}

func NewAdminAuthHandler(authService *auth.AuthService) *AdminAuthHandler {
	return &AdminAuthHandler{authService: authService}
}

func (h *AdminAuthHandler) AdminLoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}

func (h *AdminAuthHandler) AdminLoginPOST(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	userID, err := h.authService.AdminAuth(c.Request.Context(), login, password)

	if err != nil {
		log.Printf("Ошибка авторизации: %v", err)
		c.Redirect(http.StatusSeeOther, "/login")
	} else {
		isAdmin := 1
		c.SetCookie("user_id", strconv.Itoa(userID), 3600, "/", "", false, true)
		c.SetCookie("isAdmin", strconv.Itoa(isAdmin), 3600, "/", "", false, true)
		c.Redirect(http.StatusSeeOther, "/admin")
	}
}
