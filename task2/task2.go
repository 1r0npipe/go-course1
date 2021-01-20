package main

import (
	"fmt"
	"math"
	"os"
)

const MIN = 100
const MAX = 999
const MAX_ARGUMENTS = 2
func main() {
	
	var width, lenght int32
	var circleArea float64
	if len(os.Args) == MAX_ARGUMENTS {
		arg := os.Args[1]
		switch  arg {
		case "rect":
			fmt.Print("Type the width: ")
			fmt.Scanln(&width)
			fmt.Print("Input the lenght: ")
			fmt.Scanln(&lenght)
			fmt.Println("The area is", width*lenght, "of rectangle")
		case "circle":
			fmt.Print("Type the area of circle: ")
			fmt.Scanln(&circleArea)
			radius := math.Sqrt(circleArea / math.Pi)
			fmt.Printf("The diameter is: %.2f the lenght is %.2f\n", radius*2, 2*math.Pi*radius)
		case "digit":
			fmt.Print("Type any 3s digits number: ")
			fmt.Scanln(&lenght)
			if lenght < MIN || lenght > MAX {
				fmt.Println("You need to use correct number in format XXX")
				return
			}
			fmt.Printf("The hundreds: %d, then tens: %d and units: %d\n", lenght / 100, lenght % 100 / 10, lenght % 10)
		case "help":
			fmt.Println("Please use following options:")
			fmt.Println("	rect - to calculate area for the rectangle")
			fmt.Println("	circle - to calculate diameter and lenght of circle by area")
			fmt.Println("	digit (XXX) - to present hundreds, tens and units")
		default:
			fmt.Println("Use argument \"help\" after command")
		}
	} else {
		fmt.Println("No argument provided")
		return
	}
}
