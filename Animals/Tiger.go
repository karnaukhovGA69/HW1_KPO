package animals

type Tiger struct {
	food           int
	aggressiveness int
	name           string
}

func NewTiger(name_ string, food_ int, aggressiveness_ int) *Tiger {
	return &Tiger{name: name_, food: food_, aggressiveness: aggressiveness_}
}

func (t Tiger) Food() int {
	return t.food
}

func (t Tiger) Aggressiveness() int {
	return t.aggressiveness
}
