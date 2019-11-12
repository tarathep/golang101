package main

import "fmt"

func main() {
	type text = string
	var t text = "gopher is me"
	fmt.Printf("type:%T, value: %q\n", t, t)
}

/*
func main() {
	show(area)
}

//TYPE FUNCTION DECLARE
type areaFunc func(int, int) int

func area(dx, dy int) int {
	return dx * dy
}

func show(fn areaFunc) {
	fmt.Printf("area of 64x48 = %d", fn(64, 48))
}
*/
