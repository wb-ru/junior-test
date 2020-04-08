package proxy

import (
	"testing"
	"time"
)

func TestProxy(t *testing.T) {
	var number ProxyAmountOfPatients
	numberActual := number.Get()
	for i := 1; i <= 1000; i++ { // simulate a high load on the server
		if number.Get() != numberActual {
			t.Fatalf("Incorrect number of infected. Expected : %v, got : %v", numberActual, number.Get())
		}
	}
	time.Sleep(time.Second * 5)
	numberNew := number.Get()
	if numberNew == numberActual {
		t.Errorf("Information updated. Expected : %v, got : %v", numberNew, numberActual)
	}
}
