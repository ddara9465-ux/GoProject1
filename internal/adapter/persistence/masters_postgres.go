package persistence

import (
	"context"
	"log"
)

// CreateMaster создаёт нового мастера
func (r *Repository) CreateMaster(ctx context.Context, name, specialization string) error {
	_, err := r.DB.Pool.Exec(ctx, `INSERT INTO masters (name, specialization) VALUES ($1, $2)`, name, specialization)
	if err != nil {
		log.Printf("Ошибка создания мастера: %v", err)
		return err
	}
	return nil
}

// GetMasters возвращает список мастеров
func (r *Repository) GetMasters(ctx context.Context) ([]map[string]interface{}, error) {
	rows, err := r.DB.Pool.Query(ctx, `SELECT id, name, specialization FROM masters`)
	if err != nil {
		log.Printf("Ошибка получения мастеров: %v", err)
		return nil, err
	}
	defer rows.Close()

	var masters []map[string]interface{}
	for rows.Next() {
		var id int
		var name, specialization string
		if err := rows.Scan(&id, &name, &specialization); err != nil {
			log.Printf("Ошибка Scan мастера: %v", err)
			continue
		}
		masters = append(masters, map[string]interface{}{
			"ID":             id,
			"Name":           name,
			"Specialization": specialization,
		})
	}

	return masters, nil
}

// DeleteMaster удаляет мастера по ID
func (r *Repository) DeleteMaster(ctx context.Context, masterID int) error {
	_, err := r.DB.Pool.Exec(ctx, `DELETE FROM masters WHERE id = $1`, masterID)
	if err != nil {
		log.Printf("Ошибка удаления мастера: %v", err)
		return err
	}
	return nil
}
