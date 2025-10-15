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

func MinFriendLevel() int {
	return 6
}
