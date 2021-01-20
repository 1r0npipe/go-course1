package main

import (
	"fmt"	
)


func main() {
	var width, lenght int32
	fmt.Print("Inptu the width: ")
	fmt.Scanln(&width)
	fmt.Print("Input the lenght: ")
	fmt.Scanln(&lenght)
	fmt.Println("The area is", width * lenght, "of rectangle")
}