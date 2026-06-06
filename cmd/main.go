package main

import (
	myhttp "GoProject1/internal/adapter/http"
	adminHandlers "GoProject1/internal/adapter/http/handlers"
	"GoProject1/internal/adapter/persistence"
	"GoProject1/internal/application/usecases/admin"
	"GoProject1/internal/application/usecases/appointments"
	"GoProject1/internal/application/usecases/auth"
	"GoProject1/internal/application/usecases/masters"
	"GoProject1/internal/infrastructure/cache"
	"GoProject1/internal/infrastructure/db"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	ctx := context.Background()

	log.Print("Подключение к БД")
	dsn := os.Getenv("PG_DSN")
	if dsn == "" {
		log.Fatal("PG_DSN is empty")
	}

	database, err := db.NewDB(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	log.Print("Успешное подключение к БД")

	// Репозиторий
	repo := persistence.NewRepository(database)

	// Сервисы
	masterService := masters.NewMasterService(repo)
	authService := auth.NewAuthService(repo)
	adminService := admin.NewAdminService(repo)
	appointmentService := appointments.NewAppointmentService(repo)

	// Кеш
	mastersCache := cache.NewMastersCache()

	// Handlers
	aboutHandler := adminHandlers.NewAboutHandler(masterService)
	mainHandler := adminHandlers.NewMainHandler(masterService, mastersCache)
	loginHandler := adminHandlers.NewLoginHandler(authService)
	registerHandler := adminHandlers.NewRegisterHandler(authService)
	adminAuthHandler := adminHandlers.NewAdminAuthHandler(authService)
	adminPageHandler := adminHandlers.NewAdminHandler(adminService, masterService)
	adminAppointmentsHandler := adminHandlers.NewAdminAppointmentsHandler(adminService)
	appointmentsPageHandler := adminHandlers.NewAppointmentsHandler(appointmentService)

	// Передаём кеш в MastersHandler
	mastersPageHandler := adminHandlers.NewMastersHandler(masterService, mastersCache)

	// Роутер
	router := myhttp.NewRouter(
		aboutHandler,
		mainHandler,
		loginHandler,
		registerHandler,
		adminAuthHandler,
		adminPageHandler,
		adminAppointmentsHandler,
		appointmentsPageHandler,
		mastersPageHandler,
	)

	log.Print("Запуск сервера")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	router.SetupRoutes(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Print("Сервер запущен на http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Print("Получен сигнал выключения. Завершаем работу...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatal("Ошибка при выключении:", err)
	}

	log.Print("Сервер корректно завершил работу")
}
