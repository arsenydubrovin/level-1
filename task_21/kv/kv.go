package kv

import "sync"

// CSVDatabase реализует базу данных типа «ключ-значение»
type KVDatabase struct {
	Data map[int]any
	mu   sync.RWMutex
}

// NewKVDatabase создаёт экземпляр KVDatabase
func NewKVDatabase() *KVDatabase {
	return &KVDatabase{
		Data: make(map[int]any),
	}
}

// Add добавляет пару id-value в базу данных
func (db *KVDatabase) Add(id int, value any) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.Data[id] = value
}
