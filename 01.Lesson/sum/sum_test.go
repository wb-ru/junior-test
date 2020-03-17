package sum

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {

	result := twoSum([]int{1, -3, 5, 6}, -2)
	want := []int{0, 1}

	if !reflect.DeepEqual(result, want) {
		t.Error("No answer")
	}
}

func TestEmptySlice(t *testing.T) {

	if twoSum([]int{}, 10) != nil {
		t.Error("Length should be more than 0")
	}
}

func test() int {
	return 10
}
