package animals

import "fmt"

type Monkey struct {
	food         int
	friendliness int // хоть они и всеядные но по большей части для человека не представляют опасноти поэтому я решил что их можно записать как травоядных
	name         string
	number       int
}

func NewMonkey(name_ string, food_, friendliness_, number int) *Monkey {
	return &Monkey{
		name:         name_,
		food:         food_,
		friendliness: friendliness_,
		number:       number}
}

func (m Monkey) Food() int         { return m.food }
func (m Monkey) Friendliness() int { return m.friendliness }
func (m Monkey) Number() int       { return m.number }

func (m Monkey) ToString() string {
	return fmt.Sprintf("🐒 Обезьяна | Имя: %s | Еда: %d | Дружелюбность: %d | № %d",
		m.name, m.food, m.friendliness, m.number)
}
