package http

import (
	adminHandlers "GoProject1/internal/adapter/http/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	aboutHandler             *adminHandlers.AboutHandler
	mainHandler              *adminHandlers.MainHandler
	loginHandler             *adminHandlers.LoginHandler
	registerHandler          *adminHandlers.RegisterHandler
	adminAuthHandler         *adminHandlers.AdminAuthHandler
	adminHandler             *adminHandlers.AdminHandler
	adminAppointmentsHandler *adminHandlers.AdminAppointmentsHandler
	appointmentsHandler      *adminHandlers.AppointmentsHandler
	mastersHandler           *adminHandlers.MastersHandler
}

func NewRouter(
	aboutHandler *adminHandlers.AboutHandler,
	mainHandler *adminHandlers.MainHandler,
	loginHandler *adminHandlers.LoginHandler,
	registerHandler *adminHandlers.RegisterHandler,
	adminAuthHandler *adminHandlers.AdminAuthHandler,
	adminHandler *adminHandlers.AdminHandler,
	adminAppointmentsHandler *adminHandlers.AdminAppointmentsHandler,
	appointmentsHandler *adminHandlers.AppointmentsHandler,
	mastersHandler *adminHandlers.MastersHandler, // уже принимает кеш внутри
) *Router {
	return &Router{
		aboutHandler:             aboutHandler,
		mainHandler:              mainHandler,
		loginHandler:             loginHandler,
		registerHandler:          registerHandler,
		adminAuthHandler:         adminAuthHandler,
		adminHandler:             adminHandler,
		adminAppointmentsHandler: adminAppointmentsHandler,
		appointmentsHandler:      appointmentsHandler,
		mastersHandler:           mastersHandler,
	}
}

func (r *Router) SetupRoutes(engine *gin.Engine) {
	// Клиентские маршруты
	engine.GET("/login", r.loginHandler.LoginGET)
	engine.POST("/login", r.loginHandler.LoginPOST)
	engine.GET("/register", r.registerHandler.RegisterGET)
	engine.POST("/register", r.registerHandler.RegisterPOST)
	engine.GET("/", r.mainHandler.MainGET)
	engine.POST("/create-request-appointments", r.appointmentsHandler.CreateRequestAppointments)
	engine.GET("/about", r.aboutHandler.AboutGET)
	engine.POST("/logout", r.loginHandler.LogoutPOST)

	// Админские маршруты
	engine.GET("/admin", r.adminHandler.AdminMain)
	engine.POST("/admin/appointments/status", r.adminAppointmentsHandler.AdminUpdateAppointmentStatus)
	engine.POST("/admin/appointments/delete", r.adminAppointmentsHandler.AppointmentDelete)
	engine.POST("/admin/appointments/edit-details", r.adminAppointmentsHandler.AdminEditDetails)
	engine.GET("/admin-login", r.adminAuthHandler.AdminLoginGET)
	engine.POST("/admin-login", r.adminAuthHandler.AdminLoginPOST)
	engine.GET("/admin/masters", r.mastersHandler.MastersList)
	engine.POST("/admin/masters/create", r.mastersHandler.CreateMaster)
	engine.POST("/admin/masters/delete", r.mastersHandler.DeleteMaster)
}
