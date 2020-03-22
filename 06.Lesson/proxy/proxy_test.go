package proxy

import (
	"regexp"
	"testing"
)

// Just correctness test
func TestRegProxy_MatchString(t *testing.T) {
	type fields struct {
		cache map[string]*regexp.Regexp
	}
	type args struct {
		s string
		q string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Containment test", fields{}, args{"hello hello", "hello"}, true},
		{"Containment test", fields{}, args{"hello hello", "Abu Bahkr"}, false},
		{"Serious regex test", fields{}, args{"adam[23]", `^[a-z]+\[[0-9]+\]$`}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RegProxy{
				cache: tt.fields.cache,
			}
			if got := p.MatchString(tt.args.s, tt.args.q); got != tt.want {
				t.Errorf("MatchString() = %v, want %v", got, tt.want)
			}
		})
	}
}
