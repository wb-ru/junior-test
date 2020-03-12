package main

import (
	"testing"
)

func TestEmptySlice(t *testing.T) {
	got := twoSum([]int{}, 1)
	if got != nil {
		t.Errorf("Did not got nil on the empty input!")
	}
}

func TestImpossibleValue(t *testing.T) {
	got := twoSum([]int{1, 2}, -1)
	if got != nil {
		t.Errorf("Did not got nil on the impissble value")
	}
}

func TestSingleElement(t *testing.T) {
	got := twoSum([]int{1}, 1)
	want := []int{0}
	if got[0] != want[0] {
		t.Errorf("Did not got correct index on the single valued array!")
	}
}

func TestSingleElementBigArray(t *testing.T) {
	got := twoSum([]int{123123, 123123, 1, 123123, 123123123}, 1)
	want := []int{2}
	if got[0] != want[0] {
		t.Errorf("Did not got correct index on the single valued array!")
	}
}

func TestUsualArrays(t *testing.T) {
	arrays := [][]int{
		{1, 1, 2, 4, 5},
		{1, 1, 4, 1, 0},
		{-1, -10, -100},
	}
	targets := []int{
		9, 5, -110,
	}
	for i, arr := range arrays {
		target := targets[i]
		ans := twoSum(arr, target)
		if arr[ans[0]]+arr[ans[1]] != target {
			t.Errorf("You waffled with array %v and target %d", arr, target)
		}
	}
}
