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
	apps := admin.UC_getAppointments()
	mastersData := masters.UC_GetMasters()
	c.HTML(http.StatusOK, "admin.tmpl", gin.H{
		"Appointments": apps,
		"Masters":      mastersData,
	})
}
