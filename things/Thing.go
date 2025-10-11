package things

import "fmt"

type Thing struct {
	number int
	name   string
	weight float64
}

func NewThing(number int, name string, weight float64) *Thing {
	return &Thing{
		number: number,
		name:   name,
		weight: weight,
	}
}

func (t Thing) Number() int {
	return t.number
}

func (t Thing) ToString() string {
	return fmt.Sprintf("📦 Вещь | Название: %s | Вес: %.2f кг | Инв.№ %d",
		t.name, t.weight, t.number)
}
