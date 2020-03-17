package workerPool

import (
	"math/rand"
	"time"
)

func worker(jobs <-chan int, result chan<- int) {
	for j := range jobs {
		rand.New(rand.NewSource(time.Now().Unix()))
		dur := rand.Int()%3 + 1
		time.Sleep(time.Duration(dur) * time.Second)
		result <- j
	}
}
