package main

import "fmt"

func main() {
	fmt.Print("input	number:	")
	var input float64
	// %f =float
	// %d = decimal
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("output	number:", output)
}
