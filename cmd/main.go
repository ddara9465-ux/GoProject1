package main

import (
	"GoProject1/internal/adapter/http"
	"GoProject1/internal/infrastructure/db"
	"context"
	"log"
	"os"

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
	r := gin.Default()            // Создаем новый экземпляр
	r.LoadHTMLGlob("templates/*") // страницы HTML
	http.SetupRoutes(r)           // роутеры

	// Запускаем сервер на порту 8080
	r.Run(":8080")
}
