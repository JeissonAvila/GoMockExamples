package usecases

type Multiplication struct {

}

func (mul Multiplication) GetOperationResult(number1 int, number2 int) int{
	return number1 * number2
}