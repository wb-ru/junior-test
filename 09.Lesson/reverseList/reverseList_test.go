package reverseList

import (
	"testing"
)

func Test_reverseListIter(t *testing.T) {
	arrs := [][]int{
		{0, 1, 2, 3, 4, 5},
		{0},
		{},
		{1, 2},
	}

	for _, arr := range arrs {
		var l, l1, newHead *ListNode
		l = insertArr(l, arr)

		l, newHead = reverseRec(l, true)
		l1 = insertArr(l1, arr)

		l1 = reverseListIter(l1)
		if !isEqual(l1, newHead) {
			t.Errorf("Error reversing %v", arr)
		}
	}
	//type args struct {
	//	head *ListNode
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want *ListNode
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := reverseListIter(tt.args.head); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("reverseListIter() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func Test_isEqual(t *testing.T) {
	type args struct {
		other   *ListNode
		another *ListNode
	}
	var l1, l2 *ListNode
	l1 = insertArr(l1, []int{1, 2, 3})
	l2 = insertArr(l2, []int{1, 2, 3})

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"simple test", args{l1, l2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.other, tt.args.another); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
