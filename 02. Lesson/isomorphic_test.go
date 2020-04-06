package isomorphic

import "testing"

func TestIsIsomorphic(t *testing.T) {
	var tests = []struct {
		inputT string
		inputS string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"a", "b", true},
		{"aa", "bb", true},
		{"ab", "bb", false},
		{"aaaa", "bbbb", true},
		{"cat", "dog", true},
		{"caa", "aat", false},
		{"qwewq", "gthtg", true},
	}
	for _, test := range tests {
		if got := isIsomorphic(test.inputT, test.inputS); got != test.want {
			t.Errorf("isIsomorphic(%v, %v) returned %v, expected : %v", test.inputT, test.inputS, got, test.want)
		}
	}
}
