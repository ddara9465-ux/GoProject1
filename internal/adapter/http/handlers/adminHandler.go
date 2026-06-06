package http

import (
	"GoProject1/internal/application/usecases/admin"
	"GoProject1/internal/application/usecases/masters"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin — страница админки
func AdminMain(c *gin.Context) {
	//Ищем есть ли кук админа
	_, err := c.Cookie("isAdmin")
	if err != nil {
		//Если нет - направляем на главную
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	//Есть - открываем

	//  Получаем данные и показываем страницу
	ctx := c.Request.Context()
	apps := admin.UC_getAppointments(ctx)

	mastersData := masters.UC_GetMasters()

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"Appointments": apps,
		"Masters":      mastersData,
	})
}
