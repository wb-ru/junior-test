package isomorphic_strings

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Different length", args{"sdf", "d"}, false},
		{"Non-isomorphic", args{"aaa", "baa"}, false},
		{"Isomorphic simple", args{"aaa", "bbb"}, true},
		{"Isomorphic middle", args{"abbbaaacc", "abbbaaacc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
