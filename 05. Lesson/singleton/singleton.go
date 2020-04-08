package singleton

import (
	"math/rand"
	"sync"
	"time"
)

type singleton struct {
	userID int32
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
		instance = &singleton{userID: rand.Int31()}
	})
	return instance
}
