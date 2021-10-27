package usecases

const value = 2

type LocalValue struct {
}

func (local LocalValue) GetLocalValue(number int) int {
	return getLocalValue(number)
}

func getLocalValue(number int) int {
	switch number {
	case 10:
		return value-value
	case 2:
		return value*value
	default:
		return value
	}
}