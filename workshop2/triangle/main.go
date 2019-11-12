package main

import "fmt"

func main() {

	var b, h float64
CLS
	fmt.Print("---------[TRIANGLE CALCULATOR]------------\nEnter (Base Height):")
	fmt.Scanf("%f %f", &b, &h)
	fmt.Println("Area = ", (0.5 * b * h))

	fmt.Print("------------------------------------------")
}
