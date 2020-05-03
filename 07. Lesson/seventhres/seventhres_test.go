package seventhres

import (
	"reflect"
	"testing"
)

//kostyl to insert nil value in the array of int values
var NULL = -1

func createTreeNode(arr []int) *TreeNode {
	n := len(arr)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: arr[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && arr[i] != NULL {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && arr[i] != NULL {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func createIntArray(t *TreeNode) []int {
	var res []int
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				res = append(res, NULL)
			} else {
				res = append(res, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}

func TestMergeTrees(t *testing.T) {
	var tests = []struct {
		arrTree1 []int
		arrTree2 []int
		target   []int
	}{
		{[]int{1, 3, 2, 5}, []int{2, 1, 3, NULL, 4, NULL, 7}, []int{3, 4, 5, 5, 4, NULL, 7}},
	}
	for _, test := range tests {
		t1 := createTreeNode(test.arrTree1)
		t2 := createTreeNode(test.arrTree2)
		t3 := mergeTrees(t1, t2)
		t3Arr := createIntArray(t3)
		if reflect.DeepEqual(t3Arr, test.target) == false {
			t.Errorf("Test(Tree1:%v, Tree2:%v) expected to get:%v, got:%v", test.arrTree1, test.arrTree2, test.target, t3Arr)
		}
	}
}
