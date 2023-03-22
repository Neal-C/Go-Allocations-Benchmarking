package main

import (
	"bytes"
	"io"
	"testing"
)

//t -> test
//b -> benchmark

//Benchmark => hardcoded keyword
//go test -benchmem -run=^$ -bench ^BenchmarkFart github.com/Neal-C/Go-Allocations-Benchmarking
func BenchmarkFart(benchmark *testing.B){
	//initializing stuff for the code
	//not taken in count for benchmark
	count := 100;

	

	//benchmark.N => handled by go
	for i := 0; i < benchmark.N ; i++ {
		//everything inside here is the benchmarked code

		ints := []int{};

		for i := 0; i < count; i++ {
			ints = append(ints, i);
		}

	}

}
func BenchmarkPee(benchmark *testing.B){

	count := 100;

	benchmark.Run("slice not preallocated", func(b *testing.B){
		//initializing stuff for the code
		//not taken in account for benchmarking;

			//benchmark.N => handled by go
		for i := 0; i < benchmark.N ; i++ {
			//everything inside here is the benchmarked code

			ints := []int{};

			for i := 0; i < count; i++ {
				ints = append(ints, i);
			}

		}

	});

	benchmark.Run("slice preallocated as big as it will need", func(b *testing.B){
		//initializing stuff for the code
		//not taken in account for benchmarking;

			//benchmark.N => handled by go
		for i := 0; i < benchmark.N ; i++ {
			//everything inside here is the benchmarked code

			ints := make([]int, count);

			for i := 0; i < count; i++ {
				ints[i] = i;
			}

		}

	});
	



}

func BenchmarkSlicePreallocated(benchmark *testing.B){
		//initializing stuff for the code

		count := 100;

		//benchmark.N => handled by go
		for i := 0; i < benchmark.N ; i++ {
			//everything inside here is the benchmarked code
	
			ints := make([]int, count);
	
			for i := 0; i < count; i++ {
				ints[i] = i;
			}
	
		}
	
}

func Hundred100AllocsWriteBuffer(message []byte){
	buffer := new(bytes.Buffer);
	buffer.Write(message)
}

func SmartWriteBuffer(writer io.Writer, message []byte){
	writer.Write(message);
}

func BenchmarkWriteBuffer100Allocs(benchmark *testing.B){
	message := []byte("foo");
	
	for i := 0; i < benchmark.N; i++ { //internal bench handled by golang

		for i := 0; i < 100; i++ { //our logic
			Hundred100AllocsWriteBuffer(message);
		}
	}
}
func BenchmarkSmartWriteBuffer(benchmark *testing.B){
	message := []byte("foo");
	buffer := new(bytes.Buffer); //sure, 1 alloc technically
	
	for i := 0; i < benchmark.N; i++ { //internal bench handled by golang

		for i := 0; i < 100; i++ { //our logic
			//* 0 allocs
			SmartWriteBuffer(buffer, message);
		}
	}
}