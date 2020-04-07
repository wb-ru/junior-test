package workerpool

import (
	"fmt"
	"testing"
)

func benchmark(b *testing.B, amount int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		workController(amount)
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	for _, n := range []int{5, 10, 100, 1000} {
		b.Run(fmt.Sprintf("Number of workers = %v", n), func(b *testing.B) {
			benchmark(b, n)
		})
	}
}
