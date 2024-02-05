// Реализовать паттерн «адаптер» на любом примере.

// Старую базу данных CSVDatabase с методом Insert нужно заменить на новую kv.KVDatabase с методом Add. Для замены используется паттерн «адаптер».

// Target — Database
// Adaptee — kv.KVDatabase
// Adapter — KVDatabaseAdapter

package main

import (
	"github.com/arsenydubrovin/level-1/task_21/kv"
)

// App содержит зависимости приложения
type App struct {
	db Database
}

// Database описывает методы базы данных
type Database interface {
	Insert(id int, value string) error
}

// NewKVDatabaseAdapter возвращает экземпляр KVDatabaseAdapter, реализующий Database
func NewKVDatabaseAdapter(kvdb *kv.KVDatabase) Database {
	return &KVDatabaseAdapter{kvdb}
}

// KVDatabaseAdapter адаптирует kv.KVDatabase так, чтобы она была реализовывала Database
type KVDatabaseAdapter struct {
	*kv.KVDatabase // Встраивание адаптируемой структуры
}

// Insert соответствует методу Database и адаптирует метод kv.KVDatabase.Add
func (a *KVDatabaseAdapter) Insert(id int, value string) error {
	a.Add(id, value)

	return nil
}

func main() {
	// Старая база данных:
	// db := NewCSVDatabase()

	// Использование адаптера для новой базы данных
	kvdb := kv.NewKVDatabase()
	db := NewKVDatabaseAdapter(kvdb)

	// Код продолжает работать без изменения:
	a := App{db}

	a.db.Insert(1, "Константин Константинопольский")
}
