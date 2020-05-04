package ninthres

import (
	"reflect"
	"testing"
)

type list struct {
	head *ListNode
}

func initializeList() *list {
	return &list{}
}
func (l *list) addNode(value int) error {
	s := &ListNode{
		value: value,
	}
	if l.head == nil {
		l.head = s
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = s
	}
	return nil
}

//does not work for circular linked list
func ListToArray(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.value)
		if head.next == nil {
			break
		}
		head = head.next
	}
	return res
}

func TestHasCycle(t *testing.T) {
	var tests = []struct {
		nodes  []int
		target []int
	}{
		{[]int{3, 2, 0, -4}, []int{-4, 0, 2, 3}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{0}, []int{0}},
	}
	for _, test := range tests {
		list := initializeList()
		if test.nodes == nil {
			list.head = nil
		} else {
			for _, index := range test.nodes {
				list.addNode(index)
			}
		}
		reverse := reverseList(list.head)
		if got := ListToArray(reverse); reflect.DeepEqual(got, test.target) == false {
			t.Errorf("Test(List:%v) got: %v, want: %v", test.nodes, got, test.target)
		}
	}
}
