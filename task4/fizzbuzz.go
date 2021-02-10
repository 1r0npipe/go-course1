package main

import "fmt"

func fizzbuzz() {
	max := 100
	for i := 1; i <= max - 15; i = i + 15 {
		fmt.Printf("%d\n%d\nFizz\n%d\nBuzz\nFizz\n%d\n%d\nFizz\nBuzz\n%d\nFizz\n%d\n%d\nFizzBuzz\n", i, i + 1, i + 3,
		i + 6, i + 7, i + 10, i + 12, i + 13)
	}
	for i := max - 15; i <= max; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i % 3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i % 5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

}