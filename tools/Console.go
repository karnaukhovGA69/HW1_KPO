package tools

import (
	"bufio"
	"fmt"
	svc "main/service"
	"os"
	"strconv"
	"strings"
)

type Console struct {
	Zoo   svc.Zoo
	Vet   svc.VetClinic
	in    *bufio.Scanner
	items []MenuItem
}

func NewConsole(z svc.Zoo, v svc.VetClinic) *Console {
	return &Console{
		Zoo: z,
		Vet: v,
		in:  bufio.NewScanner(os.Stdin),
	}
}

func (c *Console) Run() {
	if c.items == nil {
		c.items = ReadFile("ru")
	}

	for {
		c.showMenu()
		choice := c.getUserChoice()

		if choice == len(c.items)+1 {
			fmt.Println("\n❌ Выход из программы...")
			return
		}

		if choice > 0 && choice <= len(c.items) {
			c.route(c.items[choice-1].Key)
		} else {
			fmt.Println("\n⚠️  Неверный выбор! Попробуйте снова.")
		}
	}
}

func (c *Console) showMenu() {
	fmt.Println("\n================ZOO================")
	for i, item := range c.items {
		fmt.Printf("%d) %s\n", i+1, item.Text)
	}
	fmt.Println("====================================")
}

func (c *Console) getUserChoice() int {
	fmt.Print("Выберите пункт меню: ")
	c.in.Scan()
	input := c.in.Text()
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return -1
	}
	return choice
}

func (c *Console) waitForEnter() {
	fmt.Print("\nНажмите Enter для продолжения...")
	if c.in == nil {
		fmt.Println()
		return
	}
	c.in.Scan()
}

func (c *Console) route(key string) {
	switch key {
	case "add":
		c.GetAnimal()
	case "addThing":
		c.GetThing()
	case "allAnimal":
		c.WriteAllAnimal()
		c.waitForEnter()
	case "handAnimal":
		c.WriteAllHandAnimal()
		c.waitForEnter()
	case "allZoo":
		c.WriteAllInventory()
		c.waitForEnter()
	case "changeLanguage":
		c.changeLanguage()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("\n❓ Неизвестная функция")
		c.waitForEnter()
	}
}

func (c *Console) changeLanguage() {
	fmt.Println("\n🌐 Смена языка")
	fmt.Println("Доступные языки: ru, en, ch")
	fmt.Print("Введите язык: ")

	c.in.Scan()
	language := strings.TrimSpace(c.in.Text())

	newItems := ReadFile(language)
	if newItems != nil {
		c.items = newItems
		fmt.Printf("✅ Язык успешно изменен на '%s'\n", language)
	} else {
		fmt.Printf("❌ Не удалось загрузить язык '%s'. Используется предыдущий.\n", language)
	}
	c.waitForEnter()
}
