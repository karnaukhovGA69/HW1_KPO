package things

import "fmt"

type Table struct {
	number   int
	material string
	color    string
}

func NewTable(number int, material, color string) *Table {
	return &Table{
		number:   number,
		material: material,
		color:    color,
	}
}

func (t Table) Number() int {
	return t.number
}

func (t Table) ToString() string {
	return fmt.Sprintf("🪑 Стол | Материал: %s | Цвет: %s | Инв.№ %d",
		t.material, t.color, t.number)
}
