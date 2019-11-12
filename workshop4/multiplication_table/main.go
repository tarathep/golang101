package main

import "fmt"

func main() {
	var num int
	fmt.Print("---------[MULTIPLICATION TABLE]------------\nEnter (NUM):")
	fmt.Scanf("%d", &num)
	for i := 1; i <= 12; i++ {
		fmt.Println(num, "x", i, "=", i*num)
	}
	fmt.Print("-----------------------------------------")
}
