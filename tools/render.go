package tools

import (
	"fmt"
	"strconv"
)

func (c *Console) WriteAllAnimal() {
	fmt.Print("Общее количество животных: ", len(c.Zoo.GetAllAnimals()), "\n")
	fmt.Print("Общее количество потребляемой еды: ", c.Zoo.CountFood(), "\n")
	for i, animal := range c.Zoo.GetAllAnimals() {
		fmt.Println(strconv.Itoa(i+1) + ") " + animal.ToString())
	}
}

func (c *Console) WriteAllHandAnimal() {
	fmt.Print("Общее количество животных в контактном зоопарке: ",
		len(c.Zoo.GetAllHandAnimal(MinFriendLevel())), "\n")
	handAnimals := c.Zoo.GetAllHandAnimal(MinFriendLevel())
	for i, animal := range handAnimals {
		fmt.Println(strconv.Itoa(i+1) + ") " + animal.ToString())
	}
}

func (c *Console) WriteAllInventory() {
	fmt.Println("Общий баланс зоопарка: ", len(c.Zoo.GetAllInventory()))
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
