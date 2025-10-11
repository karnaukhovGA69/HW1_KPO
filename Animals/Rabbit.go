package animals

import "fmt"

type Rabbit struct {
	name         string
	food         int
	friendliness int
	number       int
}

func NewRabbit(name string, food, friendliness, number_ int) *Rabbit {
	return &Rabbit{
		name:         name,
		food:         food,
		friendliness: friendliness,
		number:       number_,
	}
}

func (r Rabbit) Food() int         { return r.food }
func (r Rabbit) Friendliness() int { return r.friendliness }
func (r Rabbit) Number() int       { return r.number }

func (r Rabbit) ToString() string {
	return fmt.Sprintf("🐰 Кролик | Имя: %s | Еда: %d | Дружелюбность: %d | № %d",
		r.name, r.food, r.friendliness, r.number)
}
