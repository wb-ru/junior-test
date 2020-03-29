package memento

import (
	"testing"
)

func TestCareTaker_Pop(t *testing.T) {
	caretaker := new(CareTaker)

	states := [4]string{
		"1", "2", "3", "4",
	}
	orig := new(Originator)
	for _, state := range states {
		orig.state = state
		caretaker.Push(orig.GetState())
	}

	returned := [4]string{}
	index := len(returned) - 1
	for ret := caretaker.Pop(); ret != nil; {
		returned[index] = ret.state
		index--
	}

	for i, item := range states {
		if item != returned[i] {
			t.Errorf("Diff states: %s vs %s", item, returned[i])
		}
	}
}
