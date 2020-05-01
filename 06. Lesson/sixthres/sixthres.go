package sixthres

type ListNode struct {
	value int
	next  *ListNode
}

func hasCycle(head *ListNode) bool {
	first := head
	second := head
	if head == nil || head.next == nil {
		return false
	}
	for first != nil && second != nil {
		second = second.next
		first = first.next.next
		if second == first {
			return true
		}
	}
	return false
}
