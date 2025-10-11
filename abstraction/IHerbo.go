package abstraction

type IHerbo interface {
	IAlive
	IInventory
	Friendliness() int
}

/*
func NewHerbo(food_ int) *Herbo {
	return &Herbo{food: food_}
}

func (h Herbo) Food() int {
	return h.food
}
*/
