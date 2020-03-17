package factory_method

import (
	"testing"
)

func TestSong_serialize(t *testing.T) {
	type fields struct {
		author string
		text   string
	}
	type args struct {
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"JSON test",
			fields{"aasd", "asd"}, args{"json"}, "{\"Author\":\"aasd\",\"Text\":\"asd\"}",
		},

		{"XML test",
			fields{"aasd", "asd"}, args{"xml"}, "<Song><Author>aasd</Author><Text>asd</Text></Song>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Song{
				Author: tt.fields.author,
				Text:   tt.fields.text,
			}
			if got := p.serialize(tt.args.format); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
