package usecases

import (
	mock_services "MocksExamples/src/api/services/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

//var addition AdditionOperation

const (
	checkMark = "\u2713"
	failMark = "\u2717"
	mockValue = 5
)

func TestSuccessAdditionOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockObj := mock_services.NewMockLocalStorage(ctrl)
	mockObj.EXPECT().GetLocalValue(gomock.Eq(10)).Return(mockValue).AnyTimes()
	mockObj.EXPECT().GetLocalValue(gomock.Eq(8)).Return(mockValue-1).AnyTimes()
	mockObj.EXPECT().GetLocalValue(gomock.Any()).Return(mockValue+1).AnyTimes()
	var addition = AdditionOperation{
		LocalStorage: mockObj,
	}
	var tableTest = []struct{
		number1	int
		number2 int
		result int
	}{
		{
			number1: 10,
			number2: 20,
			result: 30 + mockValue,
		},
		{
			number1: 8,
			number2: 1,
			result: 9 + mockValue,
		},
		{
			number1: 5,
			number2: 10,
			result: 15 + mockValue,
		},
	}
	t.Log("Given the numbers to generate the addition operation")

	{
		for _, info := range tableTest{
			t.Logf("\t When checking first number \"%d\" and second number \"%d\"", info.number1, info.number2)
			result := addition.GetOperationResult(info.number1, info.number2)
			t.Log("\t\tShould be able to get result", checkMark)
			if result == info.result{
				t.Logf("\t\tShould have result equal \"%d\". %v", result, checkMark)
			} else {
				t.Errorf("\t\tShould have result equal \"%d\". %v %v", result, failMark, info.result)
			}
		}
	}

}
