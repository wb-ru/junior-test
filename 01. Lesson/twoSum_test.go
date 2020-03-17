package twoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		target  int
		want    [2]int
		wantErr bool
	}{
		{
			name:    "example",
			nums:    []int{2, 7, 11, 15},
			target:  9,
			want:    [2]int{0, 1},
			wantErr: false,
		},
		{
			name:    "positive case 1",
			nums:    []int{4, 2, 5, 12},
			target:  14,
			want:    [2]int{1, 3},
			wantErr: false,
		},
		{
			name:    "positive case 2",
			nums:    []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23},
			target:  24,
			want:    [2]int{0, 9},
			wantErr: false,
		},
		{
			name:    "positive case 3",
			nums:    []int{3, 2, 0, 5, 7},
			target:  5,
			want:    [2]int{0, 1},
			wantErr: false,
		},
		{
			name:    "negative case 1",
			nums:    []int{2, 7, 11, 15},
			target:  9,
			want:    [2]int{0, 2},
			wantErr: true,
		},
		{
			name:    "negative case 2",
			nums:    []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23},
			target:  24,
			want:    [2]int{0, 7},
			wantErr: true,
		},
		{
			name:    "negative case 3",
			nums:    []int{3, 2, 0, 5, 7},
			target:  7,
			want:    [2]int{2, 4},
			wantErr: true,
		},
		{
			name:    "wrong input - empty output",
			nums:    []int{1, 2, 3, 5, 7},
			target:  20,
			want:    [2]int{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) != tt.wantErr {
				t.Errorf("Result: %v, want: %v", got, tt.want)
			}
		})
	}
}
