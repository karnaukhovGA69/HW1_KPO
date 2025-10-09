package things

type Table struct {
	number int
}

func NewTable(number_ int) *Table {
	return &Table{number: number_}
}

func (t Table) Number() int {
	return t.number
}
