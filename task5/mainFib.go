package main

import (
	"fmt"
	"os"
	"strconv"
)

func getFibonacci(n uint32) uint32 {
	if n < 2 {
		return n
	}
	return getFibonacci(n-1) + getFibonacci(n-2)
}

func getFibonacciOptim(n uint32) uint32 {
	s := make(map[uint32]uint32)
	s[0] = 0
	s[1] = 1
	var getFib func(uint32) uint32

	getFib = func(num uint32) uint32 {
		if val, exist := s[num]; exist {
			return val
		}
		s[num] = getFib(num-1) + getFib(num-2)
		return s[num]
	}
	return getFib(n)
}

func main() {
	var inputLine string
	fmt.Print("Please enter the N for Fibonacci sequience: ")
	fmt.Scanf("%s", &inputLine)
	num64, err := strconv.ParseUint(inputLine, 10, 32)
	num := uint32(num64)
	if err != nil {
		fmt.Printf("You have entered wrong number, or parse is not possible for \"%s\"", inputLine)
		fmt.Println()
		os.Exit(1)
	}
	fmt.Printf("Your Fibonacci number is: %d\n", getFibonacci(num))
	fmt.Printf("Your Optimazed Fibonacci number is: %d", getFibonacciOptim(num))
}
