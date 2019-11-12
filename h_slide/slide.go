package main

import "fmt"

func main() {
	//FUNC MAKE
	x := make([]float64, 5)
	x[0] = 1
	x[1] = 2
	x[3] = 3
	/*
		... non limit size
	*/
	fmt.Println(x)

	//SLICE BASIC IS NOT FIX SIZE
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	//APPEND VALUE
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice2 := append(slice1, 9, 10)
	fmt.Println(slice1)
	fmt.Println(slice2)

	//ARR + SLICE
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 5)
	copy(slice4, slice3)
	fmt.Println(slice4)
}
