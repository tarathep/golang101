package main

import "fmt"

func main() {
	var score float64

	fmt.Print("---------[GRADE CALCULATOR]------------\nEnter (Score):")
	fmt.Scanf("%f", &score)
	fmt.Print("Your Grade => ")

	if score > 79 { //80
		fmt.Print("A")
	} else if score > 69 { //70-80
		fmt.Print("B")
	} else if score > 59 { //60 -70
		fmt.Print("C")
	} else if score > 49 { // 50-60
		fmt.Print("D")
	} else {
		fmt.Print("F")
	}
	fmt.Print("\n---------------------------------------")

}
