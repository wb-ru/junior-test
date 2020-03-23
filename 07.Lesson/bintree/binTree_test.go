package bintree

import (
	"testing"
)

func test_tree(values []int) bool {
	var t *TreeNode
	for _, val := range values {
		t = insert(t, val)
	}
	if !checkInvariant(t) {
		return false
	}
	for _, val := range values {
		if !find(t, val) {
			return false
		}
	}
	return true

}
func Test_insert(t *testing.T) {
	arrs := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{1},
	}

	for _, arr := range arrs {
		if !test_tree(arr) {
			t.Errorf("Error on %v", arr)
		}
	}
}

func Test_mergeTrees(t *testing.T) {
	arr1 := []int{1, 3, 4, 5}
	arr2 := []int{1, 2, 3}
	var t1, t2 *TreeNode
	for _, val := range arr1 {
		t1 = insert(t1, val)
	}
	for _, val := range arr2 {
		t2 = insert(t2, val)
	}
	res := mergeTrees(t1, t2)
	printLevel(res)
}
