package worker_pool

import (
	"fmt"
	"testing"
)

func Benchmark_manageWork(t *testing.B) {
	// for different numbers of workers
	rbound := 16
	for numWorkers := 1; numWorkers <= rbound; numWorkers *= 2 {
		t.Run(fmt.Sprintf("Number of workers %d", numWorkers), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				manageWork(numWorkers)
			}
		})
	}
}
