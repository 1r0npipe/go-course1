package main

import (
	"reflect"
	"testing"
)

func Test_insertSortSlice(t *testing.T) {
	testsTable := []struct {
		got  []int
		want []int
	}{
		{
			[]int{4, 6, 2, 8, 0, 3, 10, 56, 9, 1},
			[]int{0, 1, 2, 3, 4, 6, 8, 9, 10, 56},
		}, {
			[]int{4, 6, 2, 8, 0, 3, 10, 56, 9, 1},
			[]int{0, 1, 2, 3, 4, 6, 8, 9, 10, 56},
		},
	}
	for _, testCase := range testsTable {
		t.Logf("Test case agains %v, for want %v", testCase.got, testCase.want)
		if result := insertSortSlice(testCase.got); !reflect.DeepEqual(testCase.got, testCase.want) {
			t.Errorf("insertSortSlice() = %v, want %v", result, testCase.want)
		}
	}
}
