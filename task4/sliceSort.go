package main

import "fmt"

//insert sort
func insertSortSlice(array []int) []int {
	arrayLen := len(array)
	slice := make([]int, arrayLen)
	copy(slice, array)
	for i := 1; i < arrayLen; i++ {
		j := i
		for {
			if j > 0 && slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
				j--
			} else {
				break
			}
		}
	}
	return slice
}

//bubbleSort
func bubbleSortSlice(array []int) []int {
	arrayLen := len(array)
	slice := make([]int, arrayLen) 
	copy(slice, array)
	for i := 0; i < arrayLen - 1; i++ {
		for j := 0; j < arrayLen - 1; j++ {
			if (slice[j] > slice[j+1]) {
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}
	}
	return slice
}
func sliceTask() {
	testSlice := []int{44, 6, 5, -3, 6, -7, 4, 0, 99, 10}
	
	fmt.Println(testSlice)
}
