package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"log"
)

// A_RegisterTransaction регистрирует пользователя в 2 таблицы (users + clients) одной транзакцией.
func A_RegisterTransaction(password, first_name, last_name, phone, email string) (int, bool) {
	var user_id int
	login := email

	// Начинаем транзакцию.
	rt, err := db.Pool.Begin(context.Background())
	if err != nil {
		log.Fatal("1 Ошибка при создании транзакции:", err)
		return 0, true
	}

	// Если что-то пойдёт не так до Commit — откатим изменения.
	defer rt.Rollback(context.Background())

	// 1) Создаём пользователя и забираем его id через RETURNING.
	err = rt.QueryRow(
		context.Background(),
		"INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id",
		login, password,
	).Scan(&user_id)
	if err != nil {
		log.Fatal("2 Ошибка при создании транзакции:", err)
		return 0, true
	}

	// 2) Создаём клиента и привязываем к user_id.
	_, err = rt.Exec(
		context.Background(),
		"INSERT INTO clients (first_name, last_name, phone, email, user_id) VALUES ($1, $2, $3, $4, $5)",
		first_name, last_name, phone, email, user_id,
	)
	if err != nil {
		log.Fatal("3 Ошибка при создании транзакции:", err)
		return 0, true
	}

	// Фиксируем транзакцию (после этого данные реально сохранятся).
	if err = rt.Commit(context.Background()); err != nil {
		log.Printf("Не удалось зафиксировать транзакцию регистрации: %v", err)
		return 0, true
	}

	log.Println("Транзакция регистрации успешно выполнена")
	return user_id, false
}
