package binTreeSum

import "testing"

func Test_sumBetween(t *testing.T) {
	type args struct {
		arr []int
		l   int
		r   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"trivial case", args{[]int{1}, 0, 5}, 1},
		{"simple test", args{[]int{10, 5, 15, 3, 7, 18}, 7, 15}, 32},
		{"another simple test", args{[]int{10, 5, 15, 3, 7, 13, 18, 1, 6}, 6, 10}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := batchInsert(tt.args.arr)
			if got := sumBetween(tree, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("sumBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
