package tools

import (
	"fmt"
	"strconv"
)

func (c *Console) WriteAllAnimal() {
	fmt.Printf("Общее количество животных: %d\n", len(c.Zoo.GetAllAnimals()))
	fmt.Printf("Общее количество потребляемой еды: %d\n", c.Zoo.CountFood())
	for i, animal := range c.Zoo.GetAllAnimals() {
		fmt.Println(strconv.Itoa(i+1) + ") " + animal.ToString())
	}
}

func (c *Console) WriteAllHandAnimal() {
	fmt.Printf("Общее количество животных в контактном зоопарке: %d\n", len(c.Zoo.GetAllHandAnimal(MinFriendLevel())))
	handAnimals := c.Zoo.GetAllHandAnimal(MinFriendLevel())
	for i, animal := range handAnimals {
		fmt.Println(strconv.Itoa(i+1) + ") " + animal.ToString())
	}
}

func (c *Console) WriteAllInventory() {
	fmt.Printf("Общий баланс зоопарка: %d\n", len(c.Zoo.GetAllInventory()))
	for i, stuff := range c.Zoo.GetAllInventory() {
		fmt.Println(strconv.Itoa(i+1), stuff.ToString())
	}
}

// MinFriendLevel — единая точка, если решишь двигать порог (сейчас берём из service).
func MinFriendLevel() int {
	// импортировать service здесь ради одной константы не хотим —
	// этот хелпер можешь держать здесь и менять при надобности.
	// Если уже объявил service.MinFriendlinessLevel — просто верни его.
	return 6
}
