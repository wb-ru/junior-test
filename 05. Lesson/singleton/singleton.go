package singleton

import (
	"sync"
	"time"
)

type singleton struct {
	clock string
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		t := time.Now()
		instance = &singleton{clock: (t.Format(time.RFC3339))}
	})
	return instance
}
