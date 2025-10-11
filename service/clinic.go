package service

// MinFriendlinessLevel — минимальный уровень дружелюбности,
// с которого животное попадает в контактный зоопарк (строго >5 => 6).
const MinFriendlinessLevel = 6

// HealthLevel — перечесление уровней здоровья (1..5).
type HealthLevel int

const (
	Health1 HealthLevel = iota + 1 // очень плохое
	Health2                        // плохое
	Health3                        // среднее
	Health4                        // хорошее
	Health5                        // отличное
)

type VetClinic interface {
	CheckHealth(level HealthLevel) (ok bool, reason string)
}

type vetClinic struct {
	MinHealth HealthLevel
}

// NewVetClinic — по умолчанию требуем не ниже Health3.
func NewVetClinic() VetClinic {
	return &vetClinic{MinHealth: Health3}
}

func (v *vetClinic) CheckHealth(level HealthLevel) (bool, string) {
	if level < v.MinHealth {
		return false, "здоровье ниже минимального порога для приёма"
	}
	return true, ""
}
