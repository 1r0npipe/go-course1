package main

import "testing"

func Test_getFibonacci(t *testing.T) {
	testsTable := []struct {
		got  int
		want int
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
