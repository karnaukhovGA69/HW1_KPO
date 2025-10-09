package animals

type Monkey struct {
	food         int
	friendliness int // хоть они и всеядные но по большей части для человека не представляют опасноти поэтому я решил что их можно записать как травоядных
	name         string
}

func NewMonkey(name_ string, food_, friendliness_ int) *Monkey {
	return &Monkey{name: name_, food: food_, friendliness: friendliness_}
}

func (m Monkey) Food() int {
	return m.food
}

func (m Monkey) Friendliness() int {
	return m.friendliness
}
