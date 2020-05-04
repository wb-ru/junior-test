package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Errorf("Objects are not equal! first: %v,second: %v", instance1, instance2)
	}

}
