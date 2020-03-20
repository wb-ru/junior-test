package isomorphic

import ("testing")

func TestIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	test := []struct {
		name string
		args args
		want bool
	}{
		{"length", args{"aaa", "b"}, false},
		{"Not isomorphic", args{"xx", "xy"}, false},
		{"Isom", args{"aaa", "bbb"}, true},
	}
	for _, i := range test {
		t.Run(i.name, func(t *testing.T) {
			if got := isIsomorphic(i.args.s, i.args.t); got != i.want {
				t.Errorf("trouble =  %v, want %v", got, i.want)
			}
		})
	}
}
