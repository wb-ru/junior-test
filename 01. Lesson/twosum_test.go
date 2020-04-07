package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		input     []int
		targetSum int
	}{
		{[]int{2, 7, 11, 5}, 9},
		{[]int{1, 2}, 3},
		{[]int{0, -1, 1}, 0},
		{[]int{}, 5},
		{[]int{1}, 5},
		{[]int{1, 4}, 5},
		{[]int{1, 2, -5, 3, 4, 5}, 0},
	}
	for _, test := range tests {
		got := twoSum(test.input, test.targetSum)
		if len(test.input) < 2 {
			if got != nil {
				t.Errorf("If input slice length is less than 2, then twoSum() should return null")
			}
		} else {
			if got[0] == got[1] {
				t.Errorf("It is not correct to use the same indexes = %v", got[0])
				continue
			}
			if test.input[got[0]]+test.input[got[1]] != test.targetSum {
				t.Errorf("twoSum(%v, %v) returned %v, but %v + %v !=  %v", test.input, test.targetSum, got,
					test.input[got[0]], test.input[got[1]], test.targetSum)
			}
		}
	}
}
