package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

// CSVDatabase реализует базу данных, которая хранит данные в .csv
type CSVDatabase struct{}

// NewCSVDatabase создает новый экземпляр CSVDatabase
func NewCSVDatabase() *CSVDatabase {
	return new(CSVDatabase)
}

// Insert добавляет пару id-value в базу данных
func (db *CSVDatabase) Insert(id int, value string) error {
	file, err := os.OpenFile("task_21/db.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{strconv.Itoa(id), value}

	err = writer.Write(record)
	if err != nil {
		return err
	}

	return nil
}
