package handlers

import (
	"MocksExamples/src/api/entities"
	"MocksExamples/src/api/services"
	"MocksExamples/src/api/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Addition = "Suma"
	Subtraction = "Resta"
	Multiplication = "Multiplica"
	NotSupported = "Operación no permitida"
)

var (
	addition usecases.AdditionOperation
	subtraction usecases.Subtraction
	multiplication usecases.Multiplication
	localStorage services.LocalStorage
	UseCaseLocalValue usecases.LocalValue
	operation services.Operation
)

func GetOperation(c *gin.Context) {
	getOperationHandler(c)
}

func getOperationHandler(c *gin.Context) {
	var request entities.Request
	c.Bind(&request)
	var result int
	localStorage = UseCaseLocalValue
	localValue := localStorage.GetLocalValue(request.Value)
	switch request.Operation {
	case Addition:
		operation = addition
	case Subtraction:
		operation = subtraction
	case Multiplication:
		operation = multiplication
	default:
		c.JSON(http.StatusBadRequest, NotSupported)
	}
	result = operation.GetOperationResult(request.Value, localValue)
	c.JSON(http.StatusOK, result)
}
