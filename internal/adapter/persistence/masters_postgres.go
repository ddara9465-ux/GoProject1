package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"log"
)

func A_CreateMaster(name, specialization string) {
	ctx := context.Background()
	_, err := db.Pool.Exec(ctx, `INSERT INTO masters (name, specialization) VALUES ($1, $2)`, name, specialization)
	if err != nil {
		log.Fatal("Ошибка создания мастера:", err)
	}
}

func A_GetMasters() []map[string]interface{} {
	ctx := context.Background()
	rows, err := db.Pool.Query(ctx, `SELECT id, name, specialization FROM masters`)
	if err != nil {
		log.Fatal("Ошибка получения мастеров:", err)
	}
	defer rows.Close()

	var masters []map[string]interface{}
	for rows.Next() {
		var id int
		var name, specialization string
		if err := rows.Scan(&id, &name, &specialization); err != nil {
			log.Fatal("Ошибка Scan мастера:", err)
			continue
		}
		masters = append(masters, map[string]interface{}{
			"ID":             id,
			"Name":           name,
			"Specialization": specialization,
		})
	}
	return masters
}

func A_DeleteMaster(masterID int) {
	ctx := context.Background()
	_, err := db.Pool.Exec(ctx, `DELETE FROM masters WHERE id = $1`, masterID)
	if err != nil {
		log.Fatal("Ошибка удаления мастера:", err)
	}
}
