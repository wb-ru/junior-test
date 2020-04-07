package workerpool

import (
	"math/rand"
	"time"
)

var (
	workReceiverChannel chan worker
)

type worker struct {
	id         int
	workResult int
}

func work(id int) {
	workTime := rand.Int()%3 + 1
	time.Sleep(time.Second * time.Duration(workTime))
	workResult := rand.Int()%100 + 1 // worker's efficiency
	workReceiverChannel <- worker{
		id:         id,
		workResult: workResult,
	}
}

func workController(workersAmount int) []worker {
	rand.Seed(time.Now().UnixNano())
	workReceiverChannel = make(chan worker, workersAmount)
	for i := 0; i < workersAmount; i++ {
		go work(i)
	}
	workersResults := make([]worker, 0)
	for i := 0; i < workersAmount; i++ {
		workersResults = append(workersResults, <-workReceiverChannel)
	}
	return workersResults
}
