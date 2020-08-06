package main

import (
	"testing"
)

type TestCase struct {
	Coef     Coeficient
	Expected Result
	Error    error
}

var testCases = []TestCase{
	{
		Coef:     Coeficient{1, -1, -6},
		Expected: Result{3, -2},
	},
	{
		Coef:     Coeficient{1, 2, 1},
		Expected: Result{-1, -1},
	},
	{
		Coef:     Coeficient{7, -1, 2},
		Expected: Result{0, 0},
	},
}

func TestRootX2(t *testing.T) {
	for _, testCase := range testCases {
		result, _ := RootX2(testCase.Coef)
		if result != testCase.Expected {
			t.Errorf("expected: %v got: %v",
				testCase.Expected,
				result,
			)
		}
	}
}
