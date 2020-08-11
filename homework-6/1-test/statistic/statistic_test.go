package statistic

import "testing"

type TestCase struct {
	FuncName string
	Num      int
	Value    []float64
	Expected float64
}

var testCases = []TestCase{
	{
		FuncName: "Avarage",
		Num:      0,
		Value:    []float64{1, 2},
		Expected: 1.5,
	},
	{
		FuncName: "Avarage",
		Num:      1,
		Value:    []float64{3, 3, 3},
		Expected: 3,
	},
	{
		FuncName: "Avarage",
		Num:      2,
		Value:    []float64{},
		Expected: 0,
	},
	{
		FuncName: "Sum",
		Num:      0,
		Value:    []float64{1, 2},
		Expected: 3,
	},
	{
		FuncName: "Sum",
		Num:      1,
		Value:    []float64{3, 3, 3},
		Expected: 9,
	},
	{
		FuncName: "Sum",
		Num:      2,
		Value:    []float64{},
		Expected: 0,
	},
}

func TestFuncCases(t *testing.T) {
	for _, testCase := range testCases {
		var result float64
		switch testCase.FuncName {
		case "Avarage":
			result = Average(testCase.Value)
		case "Sum":
			result = Sum(testCase.Value)
		}

		if result != testCase.Expected {
			t.Errorf("for func %s case num: %d expected: %f got: %f",
				testCase.FuncName,
				testCase.Num,
				testCase.Expected,
				result,
			)
		}
	}
}

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
