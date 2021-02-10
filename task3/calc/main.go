package main

import (
	"fmt"

	"strconv"
	"strings"
)

const exitChar = "x"
const byteFloat = 32
const availableOperators = "*/+-"

// to get sum of two elements
func add(x, y float64) float64 {
	return x + y
}

// to get sub to get sum of two elementsm of two elements
func sub(x, y float64) float64 {
	return x - y
}

// to get mul of two elements
func mul(x, y float64) float64 {
	return x * y
}

// to get div of two elements
func div(x, y float64) float64 {
	return x / y
}

func main() {
	fmt.Println("This is simple calculator.")
	fmt.Print("Use x for Exit, any numbers are those operations:")
	fmt.Println("'*' - multiplicaton, '/' - division, '+' and '-' between any two numbers")
	var line string
	var operator rune
	//var result float64
	var operand1 float64
	var operand2 float64
	for {
		fmt.Print("Type operand: ")
		fmt.Scanf("%s", &line)
		if line[0] == exitChar[0] {
			return
		}
		fNumber, err := strconv.ParseFloat(line, byteFloat)
		if err != nil {
			fmt.Println("The number is not correct, please re-enter")
			continue
		}
		operand1 = fNumber
		for {
			fmt.Print("type operator: ")
			fmt.Scanf("%s", &line)
			if !(strings.Contains(availableOperators, line)) {
				continue
			} else {
				operator = rune(line[0])
				break
			}
		}
		fmt.Print("Type operand: ")
		fmt.Scanf("%s", &line)
		fNumber, err = strconv.ParseFloat(line, byteFloat)
		if err != nil {
			continue
		}
		operand2 = fNumber
		switch operator {
		case '+':
			fmt.Printf("Result = %.4f", add(operand1, operand2))
		case '-':
			fmt.Printf("Result = %.4f", sub(operand1, operand2))
		case '*':
			fmt.Printf("Result = %.4f", mul(operand1, operand2))
		case '/':
			if operand2 == 0 {
				fmt.Println("Division by zero!")
				continue
			}
			fmt.Printf("Result = %.4f", div(operand1, operand2))
		}
		fmt.Println("")
	}
}
