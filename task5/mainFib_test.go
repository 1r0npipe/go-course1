package main

import "testing"

func Test_getFibonacci(t *testing.T) {
	testsTable := []struct {
		got  uint32
		want uint32
	}{
		{11, 89}, {14, 377}, {18, 2584},
	}

	for _, testCase := range testsTable {
		result := getFibonacci(uint32(testCase.got))
		t.Logf("Current test: got %d, want %d", result, testCase.want)
		if result != uint32(testCase.want) {
			t.Errorf("The error occurs: got %d, want %d", result, testCase.want)
		}
	}
}
func Test_getFibonacciOptim(t *testing.T) {
	testsTable := []struct {
		got  uint32
		want uint32
	}{
		{4, 3}, {13, 233}, {20, 6765},
	}

	for _, testCase := range testsTable {
		result := getFibonacciOptim(uint32(testCase.got))
		t.Logf("Test against: got %d, want %d", result, testCase.want)
		if result != uint32(testCase.want) {
			t.Errorf("Error occurs: got %d, want %d", result, testCase.want)
		}
	}
}

func Benchmark_getFibonacci10(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacci(10)
	}
}

func Benchmark_getFibonacciOptim10(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacciOptim(10)
	}
}

func Benchmark_getFibonacciCycle10(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibcycle(10)
	}
}
func Benchmark_getFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacci(20)
	}	
}

func Benchmark_getFibonacciOptim20(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacciOptim(20)
	}	
}

func Benchmark_getFibonacciCycle20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibcycle(20)
	}
}

func Benchmark_getFibonacci30(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacci(30)
	}	
}

func Benchmark_getFibonacciOptim30(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacciOptim(30)
	}	
}

func Benchmark_getFibonacciCycle30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibcycle(30)
	}
}
func Benchmark_getFibonacci40(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacci(40)
	}	
}

func Benchmark_getFibonacciOptim40(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacciOptim(40)
	}	
}

func Benchmark_getFibonacciCycle40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibcycle(40)
	}
}

func Benchmark_getFibonacci50(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacci(50)
	}	
}

func Benchmark_getFibonacciOptim50(b *testing.B) {
	for i := 0; i < b.N; i++ {	
		getFibonacciOptim(50)
	}	
}

func Benchmark_getFibonacciCycle50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibcycle(50)
	}
}