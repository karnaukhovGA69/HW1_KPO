package abstraction

type IPredator interface {
	IALive
	IInventory
	Aggressiveness() int
}

/*
func NewPredator(food_ int) *Predator {
	return &Predator{food: food_}
}
*/
