package tenthres

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		arr1   []int
		arr2   []int
		target []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{[]int{1}, []int{1, 2}, []int{1}},
		{[]int{2, 2, 2, 2}, []int{1, 1}, []int{}},
	}

	for _, test := range tests {
		got := intersection(test.arr1, test.arr2)
		sort.Ints(got)
		if reflect.DeepEqual(got, test.target) == false {
			t.Errorf("Intersection(arr1:%v, arr2:%v), got:%v, target:%v", test.arr1, test.arr2, got, test.target)
		}
	}
}
