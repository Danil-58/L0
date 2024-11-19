package cache

import (
	"L0/internal/database"
	"L0/internal/models"
	"sync"
)

type Cache struct {
	db   *database.Postgres
	mu   sync.RWMutex
	data map[string]*models.OrderJSON
}

func NewCache(db *database.Postgres) *Cache {
	return &Cache{
		data: make(map[string]*models.OrderJSON),
		db:   db,
		mu:   sync.RWMutex{},
	}
}

