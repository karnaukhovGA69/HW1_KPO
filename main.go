package main

import (
	db "main/bd"
	"main/tools"
)

func main() {
	db.Connect()
	defer db.Close()
	tools.WriteConsole()

}
