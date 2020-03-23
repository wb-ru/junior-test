package linkedList

type ListNode struct {
	value int
	next  *ListNode
}

type List struct {
	value int
	head  *ListNode
	tail  *ListNode
}

func (p *List) add(value int) {
	if p.head == nil {
		p.head = new(ListNode)
		p.head.value = value
		p.tail = p.head
		return
	}
	node := new(ListNode)
	p.tail.next = node
	p.tail = node
	node.value = value
}

func (p *List) point(to int) {
	if to == -1 {
		return
	}
	where := p.head
	for i := 0; i < to; i++ {
		where = where.next
	}
	p.tail.next = where
}

func Perform(arr []int, pos int) *List {
	list := new(List)
	for _, item := range arr {
		list.add(item)
	}
	list.point(pos)
	return list
}

func Test(list *List) bool {
	where := list.tail.next
	for {
		if where == nil {
			return false
		}
		if where == list.tail {
			return true
		}
		where = where.next
	}

}
