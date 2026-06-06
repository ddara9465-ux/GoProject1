package cache

import (
	"sync"
	"time"
)

type MastersCache struct {
	data      []map[string]interface{}
	mu        sync.RWMutex
	updatedAt time.Time
}

func NewMastersCache() *MastersCache {
	return &MastersCache{}
}

func (c *MastersCache) GetMasters() []map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}

func (c *MastersCache) SetMasters(data []map[string]interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = data
	c.updatedAt = time.Now()
}

func (c *MastersCache) IsExpired() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return time.Since(c.updatedAt) > 5*time.Minute
}

// Invalidate очищает кеш (принудительное обновление)
func (c *MastersCache) Invalidate() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = nil
	c.updatedAt = time.Time{} // обнуляем время, чтобы кеш считался устаревшим
}
