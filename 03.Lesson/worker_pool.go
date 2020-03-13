package worker_pool

import (
	"math/rand"
	"time"
)

type outputPayload struct{}

func work(out chan *outputPayload) {
	d := rand.Int() % 3
	time.Sleep(time.Second * time.Duration(d))
	load := new(outputPayload)
	out <- load
}

func manageWork(numWorkers int) {
	out := make(chan *outputPayload, numWorkers)
	rand.Seed(1) // use some precise seed when debugging
	// rand.Seed(time.Now().Unix())
	for i := 0; i < numWorkers; i++ {
		work(out)
	}

	for i := 0; i < numWorkers; i++ {
		_ = <-out
		// log.Printf("Received payload #%d %v", i, &rcv)
	}
}
