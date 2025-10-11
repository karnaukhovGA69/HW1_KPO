package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type MenuItem struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

func ReadFile(language string) []MenuItem {
	filename := language + ".json"

	candidates := []string{
		filepath.Join("data", filename),             // ./data/ru.json — запуск из корня
		filepath.Join("..", "data", filename),       // ../data/ru.json — тесты из ./tests
		filepath.Join("..", "..", "data", filename), // ../../data/ru.json — на всякий
	}

	var lastErr error
	for _, path := range candidates {
		data, err := os.ReadFile(path)
		if err != nil {
			lastErr = err
			continue
		}
		textMenu := make([]MenuItem, 0, 32)
		if err := json.Unmarshal(data, &textMenu); err != nil {
			fmt.Printf("Ошибка парсинга JSON (%s): %v", path, err)
			return nil
		}
		return textMenu
	}

	fmt.Printf("Ошибка чтения файла меню (%s): %v", language, lastErr)
	return nil
}
