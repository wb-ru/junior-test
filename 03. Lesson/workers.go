package workerpool

import (
	"math/rand"
	"time"
)

func worker(results chan<- int, jobs <-chan int) {
	for i := range jobs {
		time.Sleep(time.Duration(rand.Int()%3+1) * time.Second)
		results <- i
	}
}
