package service

import "main/abstraction"

type Zoo interface {
	Admit(a abstraction.IAlive)
	AddInventory(x abstraction.IInventory)

	GetAllAnimals() []abstraction.IAlive
	GetAllHandAnimal(minFriend int) []abstraction.IHerbo
	GetAllInventory() []abstraction.IInventory

	CountFood() int
	NextInventoryNumber() int
}

type MoscowZoo struct {
	animals   []abstraction.IAlive
	hand      []abstraction.IHerbo
	inventory []abstraction.IInventory
}

func NewMoscowZoo() *MoscowZoo { return &MoscowZoo{} }

// ЛОГИКА: животное всегда идёт в общий зоопарк;
// в контактный — только если дружелюбность >= MinFriendlinessLevel.
func (z *MoscowZoo) Admit(a abstraction.IAlive) {
	z.animals = append(z.animals, a)

	if h, ok := a.(abstraction.IHerbo); ok {
		if h.Friendliness() >= MinFriendlinessLevel {
			z.hand = append(z.hand, h)
		}
	}
	if inv, ok := a.(abstraction.IInventory); ok {
		z.inventory = append(z.inventory, inv)
	}
}

func (z *MoscowZoo) AddInventory(x abstraction.IInventory) {
	z.inventory = append(z.inventory, x)
}

func (z *MoscowZoo) GetAllAnimals() []abstraction.IAlive {
	out := make([]abstraction.IAlive, len(z.animals))
	copy(out, z.animals)
	return out
}

func (z *MoscowZoo) GetAllHandAnimal(minFriend int) []abstraction.IHerbo {
	res := make([]abstraction.IHerbo, 0, len(z.hand))
	for _, h := range z.hand {
		if h.Friendliness() >= minFriend {
			res = append(res, h)
		}
	}
	return res
}

func (z *MoscowZoo) GetAllInventory() []abstraction.IInventory {
	out := make([]abstraction.IInventory, len(z.inventory))
	copy(out, z.inventory)
	return out
}

func (z *MoscowZoo) CountFood() int {
	total := 0
	for _, a := range z.animals {
		total += a.Food()
	}
	return total
}

func (z *MoscowZoo) NextInventoryNumber() int {
	return len(z.inventory) + 1
}
