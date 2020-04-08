package linkedlist

import (
	"testing"
)

func getLinkedList(listPointer *[]int, pos int) *[]ListNode {
	if listPointer == nil {
		return nil
	}
	list := *listPointer
	elements := make([]ListNode, 0)
	for _, value := range list {
		elements = append(elements, ListNode{value: value})
	}
	for i := len(list) - 1; i >= 0; i-- {
		if i == len(list)-1 {
			if pos == -1 {
				elements[i].next = nil
			} else {
				elements[i].next = &elements[pos]
			}
		} else {
			elements[i].next = &elements[i+1]
		}
	}
	return &elements
}

func TestHasCycle(t *testing.T) {
	var tests = []struct {
		input []int
		pos   int
		want  bool
	}{
		{[]int{1, 2, 3, 4}, -1, false},
		{[]int{1, 2, 3, 4}, 0, true},
		{[]int{1}, -1, false},
		{[]int{1}, 0, true},
		{nil, 0, false},
	}
	for _, test := range tests {
		list := *getLinkedList(&test.input, test.pos)
		var head *ListNode
		if len(list) == 0 {
			head = nil
		} else {
			head = &list[0]
		}
		if got := hasCycle(head); got != test.want {
			t.Errorf("hasCycle(%v, %v) returned : %v, expected : %v", test.input, test.pos, got, test.want)
		}
	}
}
