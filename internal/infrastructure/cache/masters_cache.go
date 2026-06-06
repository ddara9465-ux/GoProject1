package cache

import (
	"sync"
	"time"
)

// Структура кеша с мьютексом
type MastersCache struct {
	data      []map[string]interface{}
	mu        sync.RWMutex
	updatedAt time.Time
}

var cache = &MastersCache{}

// GetMasters - читаем из кеша (много читателей могут одновременно)
func GetMasters() []map[string]interface{} {
	cache.mu.RLock()         // Блокировка для чтения
	defer cache.mu.RUnlock() // Освобождаем
	return cache.data
}

// SetMasters - записываем в кеш (только один писатель)
func SetMasters(data []map[string]interface{}) {
	cache.mu.Lock()         // Эксклюзивная блокировка
	defer cache.mu.Unlock() // Освобождаем
	cache.data = data
	cache.updatedAt = time.Now()
}

// IsExpired - проверяем, не устарел ли кеш (5 минут)
func IsExpired() bool {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	return time.Since(cache.updatedAt) > 5*time.Minute
}
