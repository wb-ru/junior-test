package workerPool

import "testing"

func BenchmarkWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const numWorkers = 5
		jobs := make(chan int, numWorkers)
		results := make(chan int, numWorkers)

		for i := 1; i <= numWorkers; i++ {
			jobs <- i
		}

		close(jobs)

		for w := 1; w <= numWorkers; w++ {
			go worker(jobs, results)
		}

		for a := 1; a <= numWorkers; a++ {
			<-results
		}
	}
}
