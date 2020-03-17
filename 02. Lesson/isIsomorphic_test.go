package isIsomorphic

import (
	"reflect"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		t       string
		want    bool
		wantErr bool
	}{
		{
			name:    "example 1",
			s:       "egg",
			t:       "add",
			want:    true,
			wantErr: false,
		},
		{
			name:    "example 2",
			s:       "foo",
			t:       "bar",
			want:    false,
			wantErr: false,
		},
		{
			name:    "example 3",
			s:       "paper",
			t:       "title",
			want:    true,
			wantErr: false,
		},
		{
			name:    "negative case 1",
			s:       "foo",
			t:       "ego",
			want:    true,
			wantErr: true,
		},
		{
			name:    "negative case 2",
			s:       "paper",
			t:       "title",
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.s, tt.t); !reflect.DeepEqual(got, tt.want) != tt.wantErr {
				t.Errorf("Result: %v, want: %v", got, tt.want)
			}
		})
	}
}
