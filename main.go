package main

import (
	"log"
	"main/di"
)

func main() {
	app, err := di.BuildApp()
	if err != nil {
		log.Fatalf("Ошибка сборки контейнера: %v", err)
	}

	app.Run()
}
