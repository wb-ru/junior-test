package worker_pool

import (
	"math/rand"
	"sync"
	"time"
)

type outputPayload struct{}

var wg sync.WaitGroup

func work(out chan *outputPayload) {
	defer wg.Done()
	d := rand.Int() % 3
	time.Sleep(time.Second * time.Duration(d))
	load := new(outputPayload)
	out <- load
}

func manageWork(numWorkers int) {
	wg.Add(numWorkers)
	out := make(chan *outputPayload, numWorkers)
	rand.Seed(1) // use some precise seed when debugging
	// rand.Seed(time.Now().Unix())
	for i := 0; i < numWorkers; i++ {
		go work(out)
	}
	wg.Wait()
	for i := 0; i < numWorkers; i++ {
		_ = <-out
		// log.Printf("Received payload #%d %v", i, &rcv)
	}
}
