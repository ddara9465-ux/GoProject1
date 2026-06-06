package main

import (
	myhttp "GoProject1/internal/adapter/http"
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
	_ = godotenv.Load() // ищет .env в текущей рабочей директории (PWD)
	ctx := context.Background()

	// Б А З А  Д А Н Н Ы Х
	log.Print("Подключение к БД")
	dsn := os.Getenv("PG_DSN") //читаем строку коннекта к БД из файла конфигурации окружения (.env)
	if dsn == "" {
		log.Fatal("PG_DSN is empty")
	}

	err := db.InitPostgresPool(ctx, dsn) // пул соединений БД
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Успешное подключение к БД")
	}

	//З А П У С К  С А Й Т А
	log.Print("Запуск сервера")
	r := gin.Default() // Создаем новый экземпляр
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*") // страницы HTML
	myhttp.SetupRoutes(r)         // роутеры(маршруты)

	//Graceful Shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	//Запускаем сервер в горутине
	go func() {
		log.Print(" Сервер запущен на http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка: %v", err)
		}
	}()

	//Ждем сигнал выключения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Print("Получен сигнал выключения.Завершаем работу...")

	//Даем 5 секунд на завершение текущих запросов
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatal("Ошибка при выключении:", err)
	}

	log.Print("Сервер корректно завершил работу")
}
