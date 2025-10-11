package things

import "fmt"

type Computer struct {
	number int
	model  string
	ram    int
	cpu    string
}

func NewComputer(number int, model string, ram int, cpu string) *Computer {
	return &Computer{
		number: number,
		model:  model,
		ram:    ram,
		cpu:    cpu,
	}
}

func (c Computer) Number() int {
	return c.number
}

func (c Computer) ToString() string {
	return fmt.Sprintf("💻 Компьютер | Модель: %s | CPU: %s | RAM: %d ГБ | Инв.№ %d",
		c.model, c.cpu, c.ram, c.number)
}
