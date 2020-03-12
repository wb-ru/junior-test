package solution

import (
	"testing"
)

func TestEmptyArray(t *testing.T) {
	got := twoSum([]int{}, 1)
	if got != nil {
		t.Errorf("Did not got nil on the empty input!")
	}
}
