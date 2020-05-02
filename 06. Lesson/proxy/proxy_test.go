package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	var tests = []struct {
		message string
		target  string
	}{
		{"Political news", "Sorry you message has been filtered by our strong anti-political-spam-security."},
		{"Unpolitical news", "Your message has successfully reached our exchange server!"},
		{"Hi bro", "Your message has successfully reached our exchange server!"},
	}
	for _, test := range tests {

		proxy := new(proxy)

		result := proxy.Send(test.message)
		if result == test.target {
			t.Errorf("Our anti-spam filter should send:%v, but he send:%v.\n", test.target, result)
		}
	}
}
