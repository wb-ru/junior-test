package linkedlist

type ListNode struct {
	value int
	next  *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	used := make(map[*ListNode]bool)
	currentNode := head
	for {
		nextNode := currentNode.next
		if nextNode == nil {
			return false
		} else {
			if _, ok := used[nextNode]; ok {
				return true
			} else {
				used[nextNode] = true
				currentNode = nextNode
			}
		}
	}
}
