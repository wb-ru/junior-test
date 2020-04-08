package brackets

import (
	"testing"
)

func TestBracketsSequence(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"([)]", false},
		{"[]", true},
		{"{{{[[()]]}}}", true},
		{"[{][}]", false},
		{"[[", false},
		{"}}", false},
		{"[[]]{", false},
		{"][", false},
		{"{[[]]]", false},
		{"[[[]]](({}){[]})", true},
		{"", true},
	}
	for _, test := range tests {
		if got := isValid(test.input); got != test.want {
			if got {
				t.Errorf("%v is not correct bracket sequence, but you think yes", test.input)
			} else {
				t.Errorf("%v is correct bracket sequence, but you think no", test.input)
			}
		}
	}
}
