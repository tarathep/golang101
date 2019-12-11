package main

import "fmt"

func main() {

	var b, h float64

	fmt.Print("---------[TRIANGLE CALCULATOR]------------\nEnter (Base Height):")
	fmt.Scanf("%f %f", &b, &h)
	fmt.Println("Area = ", CalTriangle(b, h))

	fmt.Print("------------------------------------------")
}

func CalTriangle(base float64, height float64) float64 {
	return (0.5 * base * height)
}
