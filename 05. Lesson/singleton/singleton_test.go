package singleton

import (
	"sync"
	"sync/atomic"
	"testing"
)

var (
	wg        sync.WaitGroup
	instance1 int32
	instance2 int32
)

func setInstance1() {
	atomic.AddInt32(&instance1, GetInstance().userID)
	wg.Done()
}

func setInstance2() {
	atomic.AddInt32(&instance2, GetInstance().userID)
	wg.Done()
}

func TestGetInstance(t *testing.T) {
	wg.Add(2)
	go setInstance1()
	go setInstance2()
	wg.Wait()
	if instance1 != instance2 {
		t.Errorf("Instances are not equal. First : %v, second : %v", &instance1, &instance2)
	}
}
