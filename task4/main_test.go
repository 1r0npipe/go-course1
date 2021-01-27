package main

import (
	"testing"
)

func Test_insertSort(t *testing.T) {
	testsTable := []struct {
		got  [n]int
		want [n]int
	}{
		{
			[n]int{4, 6, 2, 8, 0, 3, 10, 56, 9, 1},
			[n]int{0, 1, 2, 3, 4, 6, 8, 9, 10, 56},
		}, {
			[n]int{4, 6, 2, 8, 0, 3, 10, 56, 9, 1},
			[n]int{0, 1, 2, 3, 4, 6, 8, 9, 10, 56},
		},
	}
	for _, testCase := range testsTable {
		result := insertSort(testCase.got)
		t.Logf("Test case got: %v, want %v", result, testCase.want)

		if result != testCase.want {
			t.Errorf("Error with testcase, got: %v, want: %v", result, testCase.want)
		}
	}
}
