package cache

import (
	"github.com/gost1k337/wb_demo/subscriber/internal/entity"
	"sync"
)

type Cache struct {
	m    sync.RWMutex
	data map[int64]entity.Order
}

func NewCache() *Cache {
	return &Cache{
		m:    sync.RWMutex{},
		data: make(map[int64]entity.Order),
	}
}

func (c *Cache) Get(id int64) (entity.Order, bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	order, ok := c.data[id]
	return order, ok
}

func (c *Cache) Set(id int64, order entity.Order) {
	c.m.Lock()
	defer c.m.Unlock()
	c.data[id] = order
}
