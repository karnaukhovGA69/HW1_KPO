package service

import (
	"main/abstraction"
)

type MoscowZoo struct {
	allAnimal      []abstraction.IALive
	allHandAnimals []abstraction.IHerbo
	allInventory   []abstraction.IInventory
}

func (mz *MoscowZoo) GetAllAnimals() []abstraction.IALive {
	if mz.allAnimal == nil {
		mz.allAnimal = make([]abstraction.IALive, 0)
	}
	return mz.allAnimal
}

func (mz *MoscowZoo) GetAllHandAnimal() []abstraction.IHerbo {
	if mz.allHandAnimals == nil {
		mz.allHandAnimals = make([]abstraction.IHerbo, 0)
	}
	return mz.allHandAnimals
}

func (mz *MoscowZoo) GetAllInventory() []abstraction.IInventory {
	if mz.allInventory == nil {
		mz.allInventory = make([]abstraction.IInventory, 0)
	}
	return mz.allInventory
}

func (mz *MoscowZoo) GetCountAllInventory() int {
	mz.GetAllInventory()
	return len(mz.allInventory)
}

func (mz *MoscowZoo) AddInventory(inventory abstraction.IInventory) {
	mz.GetAllInventory()
	mz.allInventory = append(mz.allInventory, inventory)
}

func (mz *MoscowZoo) AddAnimal(animal abstraction.IALive) {
	mz.GetAllAnimals()
	mz.allAnimal = append(mz.allAnimal, animal)
	if h, ok := animal.(abstraction.IHerbo); ok {
		mz.addHandAnimal(h)
	}
	if inv, ok := animal.(abstraction.IInventory); ok {
		mz.AddInventory(inv)
	}
}

func (mz *MoscowZoo) addHandAnimal(handAnimal abstraction.IHerbo) {
	mz.GetAllHandAnimal()
	mz.allHandAnimals = append(mz.allHandAnimals, handAnimal)
}

var Zoo = &MoscowZoo{}

func (mz *MoscowZoo) CountFood() int {
	count := 0
	animals := mz.GetAllAnimals()
	for _, animal := range animals {
		count += animal.Food()
	}
	return count
}
