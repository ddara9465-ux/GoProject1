package http

import (
	"GoProject1/internal/application/usecases/auth"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	authService *auth.AuthService
}

func NewRegisterHandler(authService *auth.AuthService) *RegisterHandler {
	return &RegisterHandler{authService: authService}
}

func (h *RegisterHandler) RegisterGET(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func (h *RegisterHandler) RegisterPOST(c *gin.Context) {
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	password := c.PostForm("password")

	userID, ok := h.authService.Register(c.Request.Context(), password, first_name, last_name, phone, email)

	if ok {
		log.Print("Успешная регистрация")
		c.Redirect(http.StatusSeeOther, "/")
		c.SetCookie("user_id", strconv.Itoa(userID), 3600, "/", "", false, true)
	} else {
		log.Print("Ошибка регистрации")
	}
}
