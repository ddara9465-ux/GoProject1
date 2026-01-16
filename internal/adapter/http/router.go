package http

import (
	http "GoProject1/internal/adapter/http/handlers"

	"github.com/gin-gonic/gin"
)

// Маршрутизация запрососв
// Запрос -вызываемая фукнция
func SetupRoutes(r *gin.Engine) {

	//log.Println("🔧 Роутер настроен")

	r.GET("/login", http.LoginGET)
	r.POST("/login", http.LoginPOST)

	r.GET("/register", http.RegisterGET)
	r.POST("/register", http.RegisterPOST)

	r.GET("/", http.MainGET)
	r.POST("/create-request-appointments", http.CreateRequestAppointments)

	r.GET("/admin", http.AdminMain)
	r.POST("/admin/appointments/status", http.AdminUpdateAppointmentStatus)

	r.POST("/admin/appointments/delete", http.AppointmentDelete)

	r.POST("/admin/appointments/edit-details", http.AdminEditDetails)

	r.GET("/admin-login", http.AdminLoginGET)
	r.POST("/admin-login", http.AdminLoginPOST)

	r.GET("/admin/masters", http.MastersList)
	r.POST("/admin/masters/create", http.CreateMaster)
	r.POST("/admin/masters/delete", http.DeleteMaster)

}
