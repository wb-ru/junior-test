package linkedList

import (
	"testing"
)

func TestList(t *testing.T) {
	type Arg struct {
		arr  []int
		pos  int
		want bool
	}

	args := []Arg{
		Arg{[]int{1, 2}, 1, true},
		Arg{[]int{1, 2}, 0, true},
		Arg{[]int{1}, -1, false},
	}

	for _, arg := range args {
		l := Perform(arg.arr, arg.pos)
		got := Test(l)
		if arg.want != got {
			t.Error("Error detecting cycle")
		}
	}
}
