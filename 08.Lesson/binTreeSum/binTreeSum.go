package binTreeSum

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	Val   int
}

func printLevel(t *TreeNode) {
	if t == nil {
		return
	}
	var q []*TreeNode
	var cur *TreeNode
	q = append(q, t)
	q = append(q, nil)

	for len(q) > 1 {
		cur = q[0]
		q = q[1:]
		if cur == nil {
			q = append(q, nil)
			fmt.Println()
		} else {
			if cur.left != nil {
				q = append(q, cur.left)
			}
			if cur.right != nil {
				q = append(q, cur.right)
			}
			fmt.Printf("%d ", cur.Val)
		}
	}
}
func checkInvariant(t *TreeNode) bool {
	var l, r, res bool
	res = true
	if t.left != nil {
		res = res && t.left.Val < t.Val
		l = checkInvariant(t.left)
		res = res && l
	}
	if t.right != nil {
		res = res && t.right.Val >= t.Val
		r = checkInvariant(t.right)
		res = res && r
	}
	return res

}

func find(t *TreeNode, val int) bool {
	if t == nil {
		return false
	}
	if t.Val == val {
		return true
	}

	if val < t.Val {
		return find(t.left, val)
	}
	return find(t.right, val)
}

func insert(t *TreeNode, val int) *TreeNode {
	if t == nil {
		t = new(TreeNode)
		t.Val = val
		return t
	}

	if val < t.Val {
		ins := insert(t.left, val)
		t.left = ins
		return t
	}
	ins := insert(t.right, val)
	t.right = ins
	return t
}

func batchInsert(arr []int) *TreeNode {
	var t *TreeNode
	for _, item := range arr {
		t = insert(t, item)
	}
	return t
}

func sumBetween(t *TreeNode, l int, r int) int {
	if t == nil {
		return 0
	}
	sum := 0
	if l <= t.Val && t.Val <= r {
		sum += t.Val
	}

	if t.left != nil {
		sum += sumBetween(t.left, l, r)
	}
	if t.right != nil {
		sum += sumBetween(t.right, l, r)
	}
	return sum
}
