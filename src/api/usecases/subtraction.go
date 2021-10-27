package usecases

type Subtraction struct {

}

func (su Subtraction) GetOperationResult(number1 int, number2 int) int {
	return number1 - number2
}