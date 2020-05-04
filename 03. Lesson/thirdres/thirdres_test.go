package thirdres

import "testing"

func BenchmarkSample(b *testing.B) {
	workers := 5
	//testing that all workers doing their jobs
	for i := 0; i < b.N; i++ {
		got := workerPool(workers)
		if got != workers {
			b.Errorf("got %v target %v", got, workers)
		}
	}
}
