package animals

type Wolf struct {
	food           int
	aggressiveness int
	name           string
}

func NewWolf(name_ string, food_, aggressiveness_ int) *Wolf {
	return &Wolf{name: name_, food: food_, aggressiveness: aggressiveness_}
}

func (w Wolf) Food() int {
	return w.food
}

func (w Wolf) Aggressiveness() int {
	return w.aggressiveness
}
