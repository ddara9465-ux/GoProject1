package http

import (
	"GoProject1/internal/application/usecases/admin"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Admin — страница админки: если нет cookie user_id, кидаем на /login, иначе проверяем права и рендерим шаблон.
func Admin(c *gin.Context) {
	// Берём user_id из cookie (если куки нет — значит пользователь не залогинен).
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		log.Print("Нет куки user_id, редирект на /login")
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// В cookie всё строкой, переводим в int, чтобы дальше работать как с id из БД.
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return
	}

	// Проверяем, что пользователь — админ
	if admin.UC_Accesadmin(userID) {
		log.Print("Открытие admin.tmpl")

		// Достаём записи из БД и отдаём их в шаблон.
		apps := admin.UC_getAppointments()
		c.HTML(http.StatusOK, "admin.tmpl", gin.H{
			"Appointments": apps,
		})
		return
	}

	// Если не админ — не пускаем в админку.
	c.Redirect(http.StatusSeeOther, "/")
}
