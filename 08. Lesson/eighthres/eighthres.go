package eighthres

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumTree(tree *TreeNode, L int, R int) (s int) {
	if tree == nil {
		return
	}
	if (L <= tree.Val) && (tree.Val <= R) {
		s += tree.Val
	}
	if tree.Left != nil {
		s += sumTree(tree.Left, L, R)
	}
	if tree.Right != nil {
		s += sumTree(tree.Right, L, R)
	}
	return
}
