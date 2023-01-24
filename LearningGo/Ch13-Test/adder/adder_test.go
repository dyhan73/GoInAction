package adder_test

import (
	"GoInAction/LearningGo/Ch13-Test/adder"
	"testing"
)

func Test_addNumber(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
