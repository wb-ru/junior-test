package work

import "testing"

func Benchmark(b *testing.B) {

	for i := 0; i < b.N; i++ {
		number := 2
		res := make(chan int, number)
		jobs := make(chan int, number)

		for j := 0; j < number; j++ {
			jobs <- j
		}

		close(jobs)

		for j := 0; j < number; j++ {
			go work(jobs, res)
		}

		for j := 0; j < number; j++ {
			<-res
		}
	}

}