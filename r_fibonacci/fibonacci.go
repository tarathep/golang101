package main

import "fmt"

// fibonacci returns a function that returns
// successive fibonacci numbers from each
// successive call
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 1000; i++ {
		fmt.Println(f())
	}
}
