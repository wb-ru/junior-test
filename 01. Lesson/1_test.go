package twoSums

import (
	"testing"
)

func TestEmpt(t *testing.T) {
	got := twoSum([]int{}, 0)
	if got != nil {
		t.Errorf("Wrong with empty massive")
	}
}

func TestImpossible(t *testing.T) {
	got := twoSum([]int{1, 2}, -1)
	if got != nil {
		t.Errorf("Wrong with impossible massive")
	}
}

func TestSingl(t *testing.T) {
	got := twoSum([]int{1}, 1)
	if got!= nil {
		t.Errorf("Wrong witn 1 element")
	}
}

func TestArrays(t *testing.T) {
	arrays := [][]int{
		{1, 2, 3, 4, 5},
		{-1, -3, -5, -7, -9},
		{1, 10, 100},
	}
	targs := []int{
		8, -4, 101,
	}
	for i, arr := range arrays {
		target := targs[i]
		ans := twoSum(arr, target)
		if arr[ans[0]]+arr[ans[1]] != target {
			t.Errorf("Problem with normal arrays")
		}
	}
}

