package main

import (
	"testing"
)
const N = 10

func Test_insert_sort(t *testing.T) {
	testsTable := []struct {
		got  [N]int{}
		want [N]int{}
	}{
		{4,6,2,8,0,3,10,56,9,1},
		{0,1,2,3,4,6,8,9,10,56},
	}{
		{4,6,2,8,0,3,10,56,9,1},
		{0,1,2,3,4,6,8,9,10,56},
	}

	for _,testCase := range testsTable {
		result := insert_sort(testCase.got)
		t.Logf("Test case got: %v, want %v", result, testCase.want)

		if result != testCase.want {
			t.Errorf("Error with testcase, got: %v, want: %v", result, testCase.want)
		}
	}
}
