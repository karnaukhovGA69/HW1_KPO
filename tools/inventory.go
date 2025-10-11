package tools

import (
	"fmt"
	"main/things"
	"strconv"
	"strings"
)

func (c *Console) GetThing() {
	fmt.Println("\n📦 Добавление предмета в инвентарь")
	fmt.Print("Выберите тип предмета (Computer, Table, Thing): ")
	c.in.Scan()
	thingType := strings.TrimSpace(c.in.Text())

	invNumber := c.Zoo.NextInventoryNumber()

	switch thingType {
	case "Computer":
		fmt.Print("Введите модель: ")
		c.in.Scan()
		model := strings.TrimSpace(c.in.Text())

		fmt.Print("Введите объем RAM (ГБ): ")
		c.in.Scan()
		ram, err := strconv.Atoi(strings.TrimSpace(c.in.Text()))
		if err != nil {
			fmt.Println("❌ Неверный ввод RAM")
			return
		}

		fmt.Print("Введите CPU: ")
		c.in.Scan()
		cpu := strings.TrimSpace(c.in.Text())

		computer := things.NewComputer(invNumber, model, ram, cpu)
		c.Zoo.AddInventory(computer)
		fmt.Println("✅ Компьютер успешно добавлен!")

	case "Table":
		fmt.Print("Введите материал: ")
		c.in.Scan()
		material := strings.TrimSpace(c.in.Text())

		fmt.Print("Введите цвет: ")
		c.in.Scan()
		color := strings.TrimSpace(c.in.Text())

		table := things.NewTable(invNumber, material, color)
		c.Zoo.AddInventory(table)
		fmt.Println("✅ Стол успешно добавлен!")

	case "Thing":
		fmt.Print("Введите название: ")
		c.in.Scan()
		name := strings.TrimSpace(c.in.Text())

		fmt.Print("Введите вес (кг): ")
		c.in.Scan()
		weight, err := strconv.ParseFloat(strings.TrimSpace(c.in.Text()), 64)
		if err != nil {
			fmt.Println("❌ Неверный ввод веса")
			return
		}

		thing := things.NewThing(invNumber, name, weight)
		c.Zoo.AddInventory(thing)
		fmt.Println("✅ Вещь успешно добавлена!")

	default:
		fmt.Println("❌ Неизвестный тип предмета. Попробуйте снова.")
	}
	c.waitForEnter()
}
