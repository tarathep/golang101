package main

import "fmt"

func main() {
	fmt.Println(func() string {
		return "Hello, Anonymous function"
	}())
}
