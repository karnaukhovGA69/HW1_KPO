package tools

import (
	"fmt"
	animals "main/Animals"
)

func (c *Console) GetAnimal() {
	fmt.Println("\n🐒 Добавление животного")
	fmt.Print("Введите тип животного (Обезьяна, Тигр, Волк, Кролик): ")
	c.in.Scan()
	animalType := c.in.Text()
	animalType = stringsTrimSpace(animalType)

	// 1) проверяем здоровье через клинику
	health, err := c.GetHealthLevel()
	if c.handleErr(err) {
		return
	}
	if ok, reason := c.Vet.CheckHealth(health); !ok {
		fmt.Println("❌ Отказ клиники:", reason)
		c.waitForEnter()
		return
	}

	// 2) собираем остальные поля
	invNumber := c.Zoo.NextInventoryNumber()

	switch animalType {
	case "Обезьяна":
		food, err := c.GetValidFood()
		if c.handleErr(err) {
			return
		}
		friend, err := c.GetFriendless()
		if c.handleErr(err) {
			return
		}
		name, err := c.GetName()
		if c.handleErr(err) {
			return
		}
		monkey := animals.NewMonkey(name, food, friend, invNumber)
		c.Zoo.Admit(monkey)
		fmt.Println("✅ Обезьяна успешно добавлена!")

	case "Тигр":
		food, err := c.GetValidFood()
		if c.handleErr(err) {
			return
		}
		name, err := c.GetName()
		if c.handleErr(err) {
			return
		}
		tiger := animals.NewTiger(name, food, invNumber)
		c.Zoo.Admit(tiger)
		fmt.Println("✅ Тигр успешно добавлен!")

	case "Волк":
		food, err := c.GetValidFood()
		if c.handleErr(err) {
			return
		}
		name, err := c.GetName()
		if c.handleErr(err) {
			return
		}
		wolf := animals.NewWolf(name, food, invNumber)
		c.Zoo.Admit(wolf)
		fmt.Println("✅ Волк успешно добавлен!")

	case "Кролик":
		food, err := c.GetValidFood()
		if c.handleErr(err) {
			return
		}
		friend, err := c.GetFriendless()
		if c.handleErr(err) {
			return
		}
		name, err := c.GetName()
		if c.handleErr(err) {
			return
		}
		rabbit := animals.NewRabbit(name, food, friend, invNumber)
		c.Zoo.Admit(rabbit)
		fmt.Println("✅ Кролик успешно добавлен!")

	default:
		fmt.Println("❌ Неизвестное животное. Попробуйте снова.")
	}

	c.waitForEnter()
}

// маленький локальный помощник, чтобы не тащить strings в импорт
func stringsTrimSpace(s string) string {
	i := 0
	j := len(s)
	for i < j && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
		i++
	}
	for j > i && (s[j-1] == ' ' || s[j-1] == '\t' || s[j-1] == '\n' || s[j-1] == '\r') {
		j--
	}
	return s[i:j]
}
