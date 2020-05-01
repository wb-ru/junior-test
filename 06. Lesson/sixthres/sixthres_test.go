package sixthres

import (
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

func (l *list) makeCycle(point int) error {
	currentNode := l.head
	cyclePoint := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	if point == -1 {
		currentNode.next = nil
	} else {
		for i := 0; i < point; i++ {
			cyclePoint = cyclePoint.next
		}
		currentNode.next = cyclePoint
	}
	return nil
}
func TestHasCycle(t *testing.T) {
	var tests = []struct {
		nodes []int
		pos   int
		want  bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1, 2, 3, 4}, -1, false},
		{[]int{0}, 0, true},
		{[]int{1}, 1, false},
		{[]int{1}, -1, false},
		{nil, 0, false},
	}
	for _, test := range tests {
		list := initializeList()
		if test.nodes == nil {
			list.head = nil
		} else {
			for _, index := range test.nodes {
				list.addNode(index)
			}
			list.makeCycle(test.pos)
		}
		if got := hasCycle(list.head); got != test.want {
			t.Errorf("Test(nodes:%v, pos:%v) got: %v, want: %v", test.nodes, test.pos, got, test.want)
		}
	}
}
