package bintree

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

func recMerge(t1 *TreeNode, t2 *TreeNode, res *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	val := 0
	if t1 != nil {
		val += t1.Val
	}
	if t2 != nil {
		val += t2.Val
	}
	res = insert(res, val)
	if t1 != nil && t1.left != nil || t2 != nil && t2.left != nil {
		var n1, n2 *TreeNode
		if t1 != nil {
			n1 = t1.left
		}
		if t2 != nil {
			n2 = t2.left
		}
		ret := recMerge(n1, n2, res.left)
		res.left = ret
	}
	if t1 != nil && t1.right != nil || t2 != nil && t2.right != nil {
		var n1, n2 *TreeNode
		if t1 != nil {
			n1 = t1.right
		}
		if t2 != nil {
			n2 = t2.right
		}
		ret := recMerge(n1, n2, res.right)
		res.right = ret
	}

	return res
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var res *TreeNode
	return recMerge(t1, t2, res)

}
