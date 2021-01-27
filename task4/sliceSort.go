package main

import "fmt"

//insert sort
func insertSortSlice(array []int) []int {
	for i := 1; i < len(array); i++ {
		j := i
		for {
			if j > 0 && array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				j--
			} else {
				break
			}
		}
	}
	return array
}

func sliceTask() {
	testArray := []int{44, 6, 5, -3, 6, -7, 4, 0, 99, 10}
	insertSortSlice(testArray)
	fmt.Println(testArray)
}
