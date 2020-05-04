package thirdres

import (
	"math/rand"
	"time"
)

func workerPool(workers int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var counter int = 0
	workerIn := make(chan int, workers)
	workerOut := make(chan bool, workers)
	for i := 0; i < workers; i++ {
		go startWorker(workerOut, workerIn)
	}
	for i := 0; i < workers; i++ {
		workerIn <- r1.Intn(3) + 1
	}
	for i := 0; i < workers; i++ {
		if <-workerOut == true {
			counter = counter + 1
		}
	}
	return counter
}

func startWorker(workerOut chan bool, workerIn chan int) {
	sleep := <-workerIn
	time.Sleep(time.Duration(sleep) * time.Second)
	workerOut <- true
}
