package main

import "fmt"

const n = 10
//insert sort
func insertSort(array [n]int) [n]int {
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

func main() {
	testArray := [n]int{4, 6, 5, 3, 6, 7, 4, 0, 9, 10}
	
	fmt.Println(insertSort(testArray))
}
