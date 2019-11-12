package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Print(a[0], a[1])
	fmt.Print(a)

	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Print(primes)

}
