package things

type Computer struct {
	number int
}

func NewComputer(number_ int) *Computer {
	return &Computer{number: number_}
}

func (c Computer) Number() int {
	return c.number
}
