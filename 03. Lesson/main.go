package lesson3

import (
	"math/rand"
	"time"
)

func worker(result chan<- int, workers <-chan int) {

	for i := range workers {
		time.Sleep(time.Duration(rand.Int()%3+1) * time.Second)
		result <- i
	}
}
