package main

import "fmt"

func main() {
	counter := factory()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func factory() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
