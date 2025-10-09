package animals

type Rabbit struct {
	food         int
	friendliness int
	name         string
}

func NewRabbit(name_ string, food_, friendliness_ int) *Rabbit {
	return &Rabbit{name: name_, food: food_, friendliness: friendliness_}
}

func (r Rabbit) Food() int {
	return r.food
}

func (r Rabbit) Friendliness() int {
	return r.friendliness
}
