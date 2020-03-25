package reverseList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert(head *ListNode, val int) *ListNode {
	old := head
	if head == nil {
		head = new(ListNode)
		head.Val = val
		return head
	}
	for head.Next != nil {
		head = head.Next
	}
	head.Next = new(ListNode)
	head.Next.Val = val
	return old
}

func insertArr(head *ListNode, arr []int) *ListNode {
	for _, item := range arr {
		head = insert(head, item)
	}
	return head
}
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func reverseListIter(head *ListNode) *ListNode {
	var prev, cur *ListNode
	if head == nil {
		return head
	}
	cur = head.Next
	prev = head
	prev.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseRec(head *ListNode, isFirst bool) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	if head.Next == nil {
		return head, head
	}
	next, tail := reverseRec(head.Next, false)
	next.Next = head
	if isFirst {
		head.Next = nil
	}
	return head, tail
}

func isEqual(other *ListNode, another *ListNode) bool {
	for {
		if other == nil && another != nil || other != nil && another == nil {
			return false
		}
		if other == nil && another == nil {
			return true
		}

		if other.Val != another.Val {
			return false
		}

		other, another = other.Next, another.Next
	}
	return true
}

//func reverseList(head *ListNode) *ListNode {
//
//}
