package mergeArrays

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Two empty arrays", args{[]int{}, []int{}}, []int{}},
		{"Two arrays without intersection", args{[]int{1, 2}, []int{3, 4}}, []int{}},
		{"Two identical arrays", args{[]int{1, 2}, []int{1, 2}}, []int{1, 2}},
		{"", args{[]int{1, 2, 3}, []int{1, 2}}, []int{1, 2}},
		{"", args{[]int{1, 2, 3, 3, 3}, []int{1, 2, 3}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
