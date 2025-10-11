package tools

import (
	"fmt"
	svc "main/service"
	"strconv"
	"strings"
)

func (c *Console) GetName() (string, error) {
	fmt.Print("Введите имя животного: ")
	c.in.Scan()
	name := strings.TrimSpace(c.in.Text())
	if name == "" {
		return "", fmt.Errorf("❌ имя не может быть пустым")
	}
	return name, nil
}

func (c *Console) GetValidFood() (int, error) {
	fmt.Print("Введите необходимое количество еды: ")
	c.in.Scan()
	input := strings.TrimSpace(c.in.Text())
	value, err := strconv.Atoi(input)
	if err != nil || value <= 0 {
		return 0, fmt.Errorf("❌ некорректное количество еды")
	}
	return value, nil
}

func (c *Console) GetFriendless() (int, error) {
	fmt.Print("Введите уровень дружелюбности (0–10): ")
	c.in.Scan()
	input := strings.TrimSpace(c.in.Text())
	value, err := strconv.Atoi(input)
	if err != nil || value < 0 || value > 10 {
		return 0, fmt.Errorf("❌ некорректный уровень дружелюбности")
	}
	return value, nil
}

func (c *Console) GetHealthLevel() (svc.HealthLevel, error) {
	fmt.Print("Введите уровень здоровья (1–5): ")
	c.in.Scan()
	input := strings.TrimSpace(c.in.Text())
	val, err := strconv.Atoi(input)
	if err != nil || val < 1 || val > 5 {
		return 0, fmt.Errorf("❌ некорректный уровень здоровья")
	}
	return svc.HealthLevel(val), nil
}

func (c *Console) handleErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		c.waitForEnter()
		return true
	}
	return false
}
