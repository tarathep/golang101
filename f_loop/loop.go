package main

import "fmt"

func main() {
	//for
	for i := 1; i < 5; i++ {
		fmt.Print(i)
	}
	//while
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Print("-", sum)

	num := []string{"a", "b", "c"}
	//foreach
	for _, i := range num {
		fmt.Print(i)
	}

}
