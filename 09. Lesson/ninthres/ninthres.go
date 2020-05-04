package ninthres

type ListNode struct {
	value int
	next  *ListNode
}

func reverseList(head *ListNode) *ListNode {
	now := head
	var prev *ListNode = nil
	for now != nil {
		temp := now.next
		now.next = prev
		prev = now
		now = temp
	}
	return prev
}
