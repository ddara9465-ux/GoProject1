package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"log"
)

func RegisterTransaction(password string, first_name string, last_name string, phone string, email string) bool {
	var user_id int
	login := email
	rt, err := db.Pool.Begin(context.Background()) // создаем транзакцию
	if err != nil {
		log.Fatal("1 Ошибка при создании транзакции:", err)
		return false
	}
	defer rt.Rollback(context.Background()) // отмена транзакции при ошибках

	err = rt.QueryRow(
		context.Background(),
		"INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id",
		login, password,
	).Scan(&user_id)

	if err != nil {
		log.Fatal("2 Ошибка при создании транзакции:", err)
		return false
	}
	_, err = rt.Exec(
		context.Background(),
		"INSERT INTO clients (first_name, last_name, phone, email, user_id) VALUES ($1, $2, $3, $4, $5)",
		first_name, last_name, phone, email, user_id,
	)
	if err != nil {
		log.Fatal("3 Ошибка при создании транзакции:", err)
		return false
	}

	err = rt.Commit(context.Background())
	if err != nil {
		log.Printf("Не удалось зафиксировать транзакцию регистрации: %v", err)
		return false
	}

	log.Println("Транзакция регистрации успешно выполнена")
	return true
}
