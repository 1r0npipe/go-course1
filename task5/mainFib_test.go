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
func Benchmark_getFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibonacci(uint32(20))
	}
}
func Benchmark_getFibonacciOptim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibonacciOptim(uint32(20))
	}
}

func Benchmark_getFibonacciCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibcycle(uint32(20))
	}
}