package animals

import "fmt"

type Wolf struct {
	name           string
	food           int
	aggressiveness int
	number         int
}

func NewWolf(name string, food, aggressiveness, number int) *Wolf {
	return &Wolf{
		name:           name,
		food:           food,
		aggressiveness: aggressiveness,
		number:         number,
	}
}

func (w Wolf) Food() int           { return w.food }
func (w Wolf) Aggressiveness() int { return w.aggressiveness }
func (w Wolf) Number() int         { return w.number }

func (w Wolf) ToString() string {
	return fmt.Sprintf("🐺 Волк | Имя: %s | Еда: %d | № %d",
		w.name, w.food, w.number)
}
