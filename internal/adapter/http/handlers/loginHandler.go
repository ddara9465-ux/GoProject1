package http

import (
	"GoProject1/internal/application/usecases/auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	authService *auth.AuthService
}

func NewLoginHandler(authService *auth.AuthService) *LoginHandler {
	return &LoginHandler{authService: authService}
}

func (h *LoginHandler) LoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *LoginHandler) LoginPOST(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	userID, ok := h.authService.Login(c.Request.Context(), login, password)

	if ok {
		c.Redirect(http.StatusSeeOther, "/")
		c.SetCookie("user_id", strconv.Itoa(userID), 3600, "/", "", false, true)
	} else {
		c.HTML(200, "login.html", gin.H{
			"Error": "Неверный логин или пароль",
		})
	}
}

func (h *LoginHandler) LogoutPOST(c *gin.Context) {
	c.SetCookie("user_id", "", -1, "/", "", false, true)
	c.SetCookie("isAdmin", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}
