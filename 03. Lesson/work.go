package work

import (
	"math/rand"
	"time"
)

func work(itog chan<- int, jobs <-chan int) {
	for j := range jobs {
		time.Sleep(time.Duration(1+rand.Int()%3)*time.Second)
		itog <- j
	}
}
