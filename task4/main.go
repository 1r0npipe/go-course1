package main

import "fmt"

const N = 10

//insert sort
func insert_sort(array [N]int) [N]int {
	for i := 1; i < len(array); i++ {
		j := i
		for {
			if j > 0 && array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				j -= 1
			} else {
				break
			}
		}
	}
	return array
}

func main() {
	test_array := [N]int{4, 6, 5, 3, 6, 7, 4, 0, 9, 10}
	
	fmt.Println(insert_sort(test_array))
}
