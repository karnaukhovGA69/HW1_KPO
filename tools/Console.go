package tools

import (
	"bufio"
	"fmt"
	animals "main/Animals"
	"main/abstraction"
	db "main/bd"
	"os"
	"strconv"
	"strings"
)

var items []MenuItem
var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func WriteConsole() {
	if items == nil {
		items = ReadFile("ru")
	}

	for {
		showMenu()
		choice := getUserChoice()

		if choice == len(items)+1 {
			fmt.Println("\n❌ Выход из программы...")
			break
		}

		if choice > 0 && choice <= len(items) {
			SearchFunction(items[choice-1].Key)
		} else {
			fmt.Println("\n⚠️  Неверный выбор! Попробуйте снова.")
		}
	}
}

func showMenu() {
	fmt.Println("\n================ZOO================")
	for i, item := range items {
		fmt.Printf("%d) %s\n", i+1, item.Text)
	}
	fmt.Println("====================================")
}

func getUserChoice() int {
	fmt.Print("Выберите пункт меню: ")
	scanner.Scan()
	input := scanner.Text()

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return -1
	}
	return choice
}

func SearchFunction(keyFunction string) {
	switch keyFunction {
	case "add":
		GetAnimal()
	case "allAnimal":
		fmt.Println("\n📋 Функция отображения всех животных еще не реализована")
		waitForEnter()
	case "handAnimal":
		fmt.Println("\n🐰 Функция отображения контактных животных еще не реализована")
		waitForEnter()
	case "allZoo":
		fmt.Println("\n📚 Функция базы данных зоопарка еще не реализована")
		waitForEnter()
	case "changeLanguage":
		changeLanguage()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("\n❓ Неизвестная функция")
		waitForEnter()
	}
}

func changeLanguage() {
	fmt.Println("\n🌐 Смена языка")
	fmt.Println("Доступные языки: ru, en, ch")
	fmt.Print("Введите язык: ")

	scanner.Scan()
	language := strings.TrimSpace(scanner.Text())

	newItems := ReadFile(language)
	if newItems != nil {
		items = newItems
		fmt.Printf("✅ Язык успешно изменен на '%s'\n", language)
	} else {
		fmt.Printf("❌ Не удалось загрузить язык '%s'. Используется предыдущий.\n", language)
	}
	waitForEnter()
}

func AddAnimal() {
	fmt.Println("\n🐒 Добавление животного")

	foodAmount := GetValidFood()
	if foodAmount <= 0 {
		fmt.Println("❌ Отмена добавления животного")
		return
	}

	fmt.Printf("✅ Успешно! Количество еды: %d\n", foodAmount)

	waitForEnter()
}

func waitForEnter() {
	fmt.Print("\nНажмите Enter для продолжения...")
	scanner.Scan()
}

func GetName() string {
	fmt.Print("Введите имя животного: ")
	scanner.Scan()
	return scanner.Text()
}

func GetValidFood() int {
	fmt.Print("Введите необходимое количество еды: ")
	scanner.Scan()
	input := scanner.Text()
	value, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || value <= 0 {
		fmt.Println("❌ Не коректные данные")
		return 0
	}
	fmt.Println("✅ Успешно!")
	waitForEnter()
	return value
}

func GetFriendless() int {
	fmt.Print("Введите уровень дружелюбности: ")
	scanner.Scan()
	input := scanner.Text()
	value, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || value < 0 || value > 10 {
		fmt.Println("❌ Не коректные данные")
		return 0
	}
	fmt.Println("✅ Успешно!")
	waitForEnter()
	return value
}

func GetAnimal() abstraction.IALive {
	fmt.Println("\n🐒 Добавление животного")
	fmt.Print("Введите тип животного (Обезьяна, Тигр, Волк, Кролик): ")
	scanner.Scan()
	switch scanner.Text() {
	case "Обезьяна":
		food := GetValidFood()
		friend := GetFriendless()
		name := GetName()
		db.SaveAnimal(name, "Обезьяна", food, friend, 0)
		return animals.NewMonkey(name, food, friend)

	case "Тигр":
		food := GetValidFood()
		name := GetName()
		db.SaveAnimal(name, "Тигр", food, 0, 0)
		return animals.NewTiger(name, food, 0)

	case "Волк":
		food := GetValidFood()
		name := GetName()
		db.SaveAnimal(name, "Волк", food, 0, 0)
		return animals.NewWolf(name, food, 0)

	case "Кролик":
		food := GetValidFood()
		friend := GetFriendless()
		name := GetName()
		db.SaveAnimal(name, "Кролик", food, friend, 0)
		return animals.NewRabbit(name, food, friend)

	default:
		fmt.Println("Неизвестное животное.")
		return nil
	}
}
