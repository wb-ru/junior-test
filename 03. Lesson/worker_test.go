package workerpool

import "testing"

func Benchmark(b *testing.B) {

	for i := 0; i < b.N; i++ {
		numOfWorkers := 4
		results := make(chan int, numOfWorkers)
		jobs := make(chan int, numOfWorkers)

		for j := 0; j < numOfWorkers; j++ {
			jobs <- j
		}

		close(jobs)

		for j := 0; j < numOfWorkers; j++ {
			go work(jobs, results)
		}

		for j := 0; j < numOfWorkers; j++ {
			<-results
		}
	}

}
