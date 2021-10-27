package usecases

import "MocksExamples/src/api/services"

type AdditionOperation struct {
	LocalStorage services.LocalStorage
}

func (addition AdditionOperation) GetOperationResult(number1 int, number2 int) int{
	localValue := addition.LocalStorage.GetLocalValue(number1)
	return number1 + number2 + localValue
}