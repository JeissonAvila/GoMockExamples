package services

type Operation interface {
	GetOperationResult(inputValue int, systemValue int) int
}