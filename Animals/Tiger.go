package animals

import "fmt"

type Tiger struct {
	name           string
	food           int
	aggressiveness int
	number         int
}

func NewTiger(name string, food, aggressiveness, number int) *Tiger {
	return &Tiger{
		name:           name,
		food:           food,
		aggressiveness: aggressiveness,
		number:         number,
	}
}

// --- Реализация интерфейсов ---
func (t Tiger) Food() int           { return t.food }
func (t Tiger) Aggressiveness() int { return t.aggressiveness }
func (t Tiger) Number() int         { return t.number }

func (t Tiger) ToString() string {
	return fmt.Sprintf("🐯 Тигр | Имя: %s | Еда: %d | № %d",
		t.name, t.food, t.number)
}
