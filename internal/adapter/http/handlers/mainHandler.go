package http

import (
	"GoProject1/internal/application/usecases/masters"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// показать главную страницу, если уже прошли аутентификацию
func MainGET(c *gin.Context) {
	//проверяем кук входа, если нет - перенаправляет на /login
	cookie, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	} else {
		log.Print("Coockie:", cookie)

		mastersData := masters.UC_GetMasters()
		c.HTML(http.StatusOK, "main.html", gin.H{
			"masters": mastersData,
		})
	}

}
