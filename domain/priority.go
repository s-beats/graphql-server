package domain

type Priority string

const (
	high   Priority = "HIGH"
	middle Priority = "MIDDLE"
	low    Priority = "LOW"
)

func (p Priority) Value() string {
	return string(p)
}

func NewPriority(val string) Priority {
	switch val {
	case string(high):
		return high
	case string(middle):
		return middle
	case string(low):
		return low
	default:
		panic("invalid priority value")
	}
}
