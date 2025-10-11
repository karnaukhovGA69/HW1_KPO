package tools

import (
	"bufio"
	"fmt"
	animals "main/Animals"
	. "main/service"
	"main/things"
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

func waitForEnter() {
	fmt.Print("\nНажмите Enter для продолжения...")
	scanner.Scan()
}

func SearchFunction(keyFunction string) {
	switch keyFunction {
	case "add":
		GetAnimal()
	case "addThing":
		GetThing()

	case "allAnimal":
		WriteAllAnimal()
		waitForEnter()
	case "handAnimal":
		WriteAllHandAnimal()
		waitForEnter()
	case "allZoo":
		WriteAllInventory()
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

func GetName() (string, error) {
	fmt.Print("Введите имя животного: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	if name == "" {
		return "", fmt.Errorf("❌ имя не может быть пустым")
	}
	return name, nil
}

func GetValidFood() (int, error) {
	fmt.Print("Введите необходимое количество еды: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	value, err := strconv.Atoi(input)
	if err != nil || value <= 0 {
		return 0, fmt.Errorf("❌ некорректное количество еды")
	}
	return value, nil
}

func GetFriendless() (int, error) {
	fmt.Print("Введите уровень дружелюбности (0–10): ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	value, err := strconv.Atoi(input)
	if err != nil || value < 0 || value > 10 {
		return 0, fmt.Errorf("❌ некорректный уровень дружелюбности")
	}
	return value, nil
}

func GetAnimal() {
	fmt.Println("\n🐒 Добавление животного")
	fmt.Print("Введите тип животного (Обезьяна, Тигр, Волк, Кролик): ")
	scanner.Scan()
	animalType := strings.TrimSpace(scanner.Text())
	switch animalType {
	case "Обезьяна":
		food, err := GetValidFood()
		if handleErr(err) {
			return
		}

		friend, err := GetFriendless()
		if handleErr(err) {
			return
		}

		name, err := GetName()
		if handleErr(err) {
			return
		}

		monkey := animals.NewMonkey(name, food, friend, Zoo.GetCountAllInventory()+1)
		Zoo.AddAnimal(monkey)
		fmt.Println("✅ Обезьяна успешно добавлена!")

	case "Тигр":
		food, err := GetValidFood()
		if handleErr(err) {
			return
		}

		name, err := GetName()
		if handleErr(err) {
			return
		}

		tiger := animals.NewTiger(name, food, 0, Zoo.GetCountAllInventory()+1)
		Zoo.AddAnimal(tiger)
		fmt.Println("✅ Тигр успешно добавлен!")

	case "Волк":
		food, err := GetValidFood()
		if handleErr(err) {
			return
		}

		name, err := GetName()
		if handleErr(err) {
			return
		}

		wolf := animals.NewWolf(name, food, 0, Zoo.GetCountAllInventory()+1)
		Zoo.AddAnimal(wolf)
		fmt.Println("✅ Волк успешно добавлен!")

	case "Кролик":
		food, err := GetValidFood()
		if handleErr(err) {
			return
		}

		friend, err := GetFriendless()
		if handleErr(err) {
			return
		}

		name, err := GetName()
		if handleErr(err) {
			return
		}

		rabbit := animals.NewRabbit(name, food, friend, Zoo.GetCountAllInventory()+1)
		Zoo.AddAnimal(rabbit)
		fmt.Println("✅ Кролик успешно добавлен!")

	default:
		fmt.Println("❌ Неизвестное животное. Попробуйте снова.")
	}

	waitForEnter()
}

func handleErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		waitForEnter()
		return true
	}
	return false
}

func WriteAllAnimal() {
	fmt.Print("Общее количество животных: ", len(Zoo.GetAllAnimals()), "\n")
	fmt.Print("Общее количество потребляемой еды: ", Zoo.CountFood(), "\n")
	for i, animal := range Zoo.GetAllAnimals() {
		fmt.Println(strconv.Itoa(i+1) + ") " + animal.ToString())
	}

}
func WriteAllHandAnimal() {
	fmt.Print("Общее количество животных в контактном зоопарке: ", len(Zoo.GetAllHandAnimal()), "\n")
	handAnimals := Zoo.GetAllHandAnimal()

	for i, animal := range handAnimals {
		if animal.Friendliness() > 5 {
			fmt.Println(strconv.Itoa(i+1) + ") " + animal.ToString())
		}
	}

}

func GetThing() {
	fmt.Println("\n📦 Добавление предмета в инвентарь")
	fmt.Print("Выберите тип предмета (Computer, Table, Thing): ")
	scanner.Scan()
	thingType := strings.TrimSpace(scanner.Text())

	invNumber := Zoo.GetCountAllInventory() + 1

	switch thingType {
	case "Computer":
		fmt.Print("Введите модель: ")
		scanner.Scan()
		model := strings.TrimSpace(scanner.Text())

		fmt.Print("Введите объем RAM (ГБ): ")
		scanner.Scan()
		ram, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("❌ Неверный ввод RAM")
			return
		}

		fmt.Print("Введите CPU: ")
		scanner.Scan()
		cpu := strings.TrimSpace(scanner.Text())

		computer := things.NewComputer(invNumber, model, ram, cpu)
		Zoo.AddInventory(computer)
		fmt.Println("✅ Компьютер успешно добавлен!")

	case "Table":
		fmt.Print("Введите материал: ")
		scanner.Scan()
		material := strings.TrimSpace(scanner.Text())

		fmt.Print("Введите цвет: ")
		scanner.Scan()
		color := strings.TrimSpace(scanner.Text())

		table := things.NewTable(invNumber, material, color)
		Zoo.AddInventory(table)
		fmt.Println("✅ Стол успешно добавлен!")

	case "Thing":
		fmt.Print("Введите название: ")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())

		fmt.Print("Введите вес (кг): ")
		scanner.Scan()
		weight, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Println("❌ Неверный ввод веса")
			return
		}

		thing := things.NewThing(invNumber, name, weight)
		Zoo.AddInventory(thing)
		fmt.Println("✅ Вещь успешно добавлена!")

	default:
		fmt.Println("❌ Неизвестный тип предмета. Попробуйте снова.")
	}
	waitForEnter()
}

func WriteAllInventory() {
	fmt.Println("Общий баланс зоопарка: ", len(Zoo.GetAllInventory()))
	for i, stuff := range Zoo.GetAllInventory() {
		fmt.Println(strconv.Itoa(i+1), stuff.ToString())
	}
}
