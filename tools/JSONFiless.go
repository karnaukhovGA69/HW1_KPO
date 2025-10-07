package tools

import (
	"encoding/json"
	"log"
	"os"
)

type MenuItem struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

func ReadFile(language string) []MenuItem {
	path := "data/" + language + ".json"
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		return nil
	}
	textMenu := make([]MenuItem, 0, 1000000)
	if err := json.Unmarshal(data, &textMenu); err != nil {
		log.Printf("Ошибка парсинга JSON: %v", err)
		return nil
	}

	return textMenu
}
