package benchmarks

import (
	"sync"
	"testing"
)

func BenchmarkAddOne(b *testing.B) {
	b.ReportAllocs()
	var count int
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
}

func BenchmarkAddMany(b *testing.B) {
	b.ReportAllocs()
	var count int
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			count++
		}()
	}
}

//to run benchmarks you need to make sure your file ends with _test.go
//create a function starts with bench(autocomplete can help you) and use the b for measure stuff
//when you want to run for benchmarking you should go to the terminal and go to the directory where you put your benchmark files
//then write <go test -bench="files name"> instead of the files name you can write . for all files inside
