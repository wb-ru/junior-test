package eighthres

import (
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
		arrTree []int
		L       int
		R       int
		target  int
	}{
		{[]int{10, 5, 15, 3, 7, NULL, 18}, 7, 15, 32},
		{[]int{10, 5, 15, 3, 7, 13, 18, 1, NULL, 6}, 6, 10, 23},
	}
	for _, test := range tests {
		tree := createTreeNode(test.arrTree)
		tSum := sumTree(tree, test.L, test.R)
		if tSum != test.target {
			t.Errorf("Test(Tree:%v, L:%v, R:%v) expected to get:%v, got:%v", test.arrTree, test.L, test.R, test.target, tSum)
		}
	}
}
