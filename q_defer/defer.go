package main

import "fmt"

func main() {
	defer fmt.Println("defer first")
	defer fmt.Println("defer second")
	defer fmt.Println("defer third")

	fmt.Println("Hello, Gophers")
}
