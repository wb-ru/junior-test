package isomorphic

import "testing"

func TestValues(t *testing.T) {

	tests := []struct {
		firstString  string
		secondString string
		want         bool
	}{
		{"bob", "kek", true},
		{"paper", "lobst", false},
		{"aaaccckkk", "mmawerttf", false},
	}

	for i, test := range tests {
		if got := isIsomorphic(test.firstString, test.secondString); got != test.want {
			t.Error("Fail on test number: {}", i)
		}
	}
}
