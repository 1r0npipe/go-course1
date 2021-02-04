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
	
	var width, lenght uint32
	var circleArea float64
	if len(os.Args) == MAX_ARGUMENTS {
		arg := os.Args[1]
		switch  arg {
		case "rect":
			fmt.Print("Type the width: ")
			_, err := fmt.Scanln(&width)
			if err != nil {
				fmt.Println("Something goes wrong")
			}
			fmt.Print("Input the lenght: ")
			_, err = fmt.Scanln(&lenght)
			if err != nil {
				fmt.Println("Something goes wrong")
			}
			fmt.Println("The area is", width*lenght, "of rectangle")
		case "circle":
			fmt.Print("Type the area of circle: ")
			_, err := fmt.Scanln(&circleArea)
			if err != nil {
				fmt.Println("Something goes wrong")
			}
			radius := math.Sqrt(circleArea / math.Pi)
			fmt.Printf("The diameter is: %.2f the lenght is %.2f\n", radius*2, 2*math.Pi*radius)
		case "digit":
			fmt.Print("Type any 3s digits number: ")
			_, err := fmt.Scanln(&lenght)
			if err != nil {
				fmt.Println("Something goes wrong")
			}
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
