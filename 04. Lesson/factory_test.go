package factory

import (
	"reflect"
	"testing"
)

func TestLogistic(t *testing.T) {
	tests := []struct {
		name   string
		method string
		want   string
	}{
		{
			name:   "case 1",
			method: "Ground",
			want:   "byGround",
		},
		{
			name:   "positive case 2",
			method: "Air",
			want:   "byAir",
		},
		{
			name:   "positive case 3",
			method: "Sea",
			want:   "bySea",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reflect.TypeOf(logistic(tt.method)).Name(); !reflect.DeepEqual(got, tt.want) {
				t.Error("Result:", got, " want: ", tt.want)
			}
		})
	}
}

func TestPutGet(t *testing.T) {
	tests := []struct {
		name   string
		method string
		id     int
		From   string
		To     string
		want   struct {
			from string
			to   string
		}
		wantErr bool
	}{
		{
			name:   "case 4",
			method: "Air",
			id:     1,
			From:   "Moscow",
			To:     "Saint Petersburg",
			want: struct {
				from string
				to   string
			}{from: "Moscow", to: "Saint Petersburg"},
			wantErr: false,
		},
		{
			name:   "case 5",
			method: "Ground",
			id:     1,
			From:   "London",
			To:     "Paris",
			want: struct {
				from string
				to   string
			}{from: "London", to: "Paris"},
			wantErr: false,
		},
		{
			name:   "case 6",
			method: "Sea",
			id:     1,
			From:   "Amsterdam",
			To:     "New York",
			want: struct {
				from string
				to   string
			}{from: "Amsterdam", to: "New York"},
			wantErr: false,
		},
		{
			name:   "case 7",
			method: "Air",
			id:     1,
			From:   "Smolensk",
			To:     "Pskov",
			want: struct {
				from string
				to   string
			}{from: "Moscow", to: "Saint Petersburg"},
			wantErr: true,
		},
		{
			name:   "case 8",
			method: "Ground",
			id:     1,
			From:   "Warsaw",
			To:     "Paris",
			want: struct {
				from string
				to   string
			}{from: "London", to: "Paris"},
			wantErr: true,
		},
		{
			name:   "case 6",
			method: "Sea",
			id:     1,
			From:   "Amsterdam",
			To:     "New York",
			want: struct {
				from string
				to   string
			}{from: "Berlin", to: "Vienna"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := logistic(tt.method)
			data.putInfo(tt.id, tt.From, tt.To)
			from, to := data.getInfo(tt.id)
			got := struct {
				from string
				to   string
			}{
				from: from,
				to:   to,
			}
			if !reflect.DeepEqual(got, tt.want) != tt.wantErr {
				t.Errorf("Result: %v, want: %v", got, tt.want)
			}
		})
	}
}
