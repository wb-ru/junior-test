package fifthres

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		bracks string
		want   bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"]", false},
		{"[", false},
		{"][", false},
		{"[[]]", true},
		{"[[]", false},
		{"[]}", false},
	}
	for _, test := range tests {
		got := isValid(test.bracks)
		if got != test.want {
			t.Errorf("%v result is: %v, you got: %v", test.bracks, test.want, got)
		}
	}
}
