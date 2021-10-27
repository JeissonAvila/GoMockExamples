package usecases

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var multiplication Multiplication

func TestSuccessMultiplicationOperation(t *testing.T) {
	var tableTest = []struct{
		number1	int
		number2 int
		result int
	}{
		{
			number1: 10,
			number2: 20,
			result: 200,
		},
		{
			number1: 8,
			number2: 1,
			result: 8,
		},
		{
			number1: 5,
			number2: 10,
			result: 50,
		},
	}
	t.Log("Given the numbers to generate the multiplication operation")

	{
		for _, info := range tableTest{
			t.Logf("\t When checking first number \"%d\" and second number \"%d\"", info.number1, info.number2)
			result := multiplication.GetOperationResult(info.number1, info.number2)
			t.Log("\t\tShould be able to get result", checkMark)
			if assert.Equal(t, info.result, result){
				t.Logf("\t\tShould have result equal \"%d\". %v", result, checkMark)
			} else {
				t.Errorf("\t\tShould have result equal \"%d\". %v %v", result, failMark, info.result)
			}
		}
	}

}
