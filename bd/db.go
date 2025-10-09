package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *sql.DB

func Connect() {
	var err error
	Conn, err = sql.Open("sqlite3", "./zoo.db")
	if err != nil {
		log.Fatalf("Ошибка открытия БД: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS animals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		type TEXT NOT NULL,
		food INTEGER NOT NULL,
		friendliness INTEGER,
		aggressiveness INTEGER
	);`

	_, err = Conn.Exec(createTable)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}

	fmt.Println("✅ SQLite база успешно подключена!")
}

func Close() {
	if Conn != nil {
		Conn.Close()
	}
}

func SaveAnimal(name, animalType string, food, friendliness, aggressiveness int) error {
	_, err := Conn.Exec(
		`INSERT INTO animals (name, type, food, friendliness, aggressiveness)
		 VALUES (?, ?, ?, ?, ?)`,
		name, animalType, food, friendliness, aggressiveness)
	if err != nil {
		return fmt.Errorf("ошибка сохранения животного: %v", err)
	}
	fmt.Printf("✅ Животное '%s' успешно сохранено в SQLite!\n", name)
	return nil
}
