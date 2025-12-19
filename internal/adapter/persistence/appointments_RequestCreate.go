package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"log"
)

func A_CreateRequestAppointments(date string, employee string, procedure string, notes string) {
	client_id := 1
	status := "Запрос звонка"
	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO appointments (client_id, procedure, employee, appointment_date, status, notes) VALUES ($1,$2,$3,$4,$5,$6)",
		client_id, procedure, employee, date, status, notes,
	)
	if err != nil {
		log.Fatal("Ошибка при создании запроса записи:", err)
	}

}
