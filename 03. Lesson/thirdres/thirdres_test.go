package thirdres

import "testing"

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		workerPool(5)
	}
}
