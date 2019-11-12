package main

import "fmt"

func main() {
	show(area)
}

func area(dx, dy int) int {
	return dx * dy
}

func show(fn func(int, int) int) {
	fmt.Printf("area of 64x48 = %d", fn(64, 48))
}
