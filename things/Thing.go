package things

type Thing struct {
	number int
}

func NewThing(number_ int) *Thing {
	return &Thing{number: number_}
}

func (t Thing) Number() int {
	return t.number
}
