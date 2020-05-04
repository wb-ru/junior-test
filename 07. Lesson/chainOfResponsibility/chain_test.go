package chainofresponsibility

import (
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	hCritical := new(HandlerCritical)
	hShops := new(HandlerShops)
	hWork := new(HandlerWork)
	hWork.next = hShops
	hShops.next = hCritical
	var tests = []struct {
		message string
		target  string
	}{
		{"Monitoring", "Message sent to critical folder."},
		{"Shop", "Message sent to shop folder."},
		{"Job", "Message sent to job folder."},
		{"Spam", ""},
	}
	for _, test := range tests {
		result := hWork.SendRequest(test.message)
		if result != test.target {
			t.Errorf("Expect result to be:%v, but got:%v", test.target, result)
		}
	}
}
