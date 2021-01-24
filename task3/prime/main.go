package main

import (
	"fmt"
	"math"
)

func isPrime(a int) bool{
	if a <= 2 { return true }
	for j := 2; j <= int(math.Sqrt(float64(a))) ; j++ {
		if a % j == 0 { 
			return false
		}
	}
	return true
}


func main() {
	var num int
	fmt.Print("Type any number, to find all primes to that:")
	fmt.Scanln(&num)
	fmt.Printf("The list of primes between 1 and %d\n", num)
	for i := 1; i < num; i++ {
		if isPrime(i) { fmt.Printf(" %d",i)}
	}
	fmt.Println()
}
