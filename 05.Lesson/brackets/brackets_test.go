package brackets

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"tmixed", args{"[)"}, false},
		{"empty string", args{""}, true},
		{"one type", args{"()()()()"}, true},
		{"one type, false", args{"())"}, false},
		{"two types", args{"[]()"}, true},
		{"two types bad", args{"[]({)"}, false},
		{"three types", args{"[](){}"}, true},
		{"folding", args{"((((()))))"}, true},
		{" false folding", args{"(((((()))))"}, false},
		{" false folding", args{"(((((({[]})))))"}, false},
		{" folding", args{"[]{}([{}])[](())"}, true},
		{" mixed folding", args{"[(])"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
