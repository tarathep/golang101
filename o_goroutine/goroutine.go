package main

import (
	"fmt"
	"time"
)

func printout(i int) {
	fmt.Print(i)
}

func main() {
	total := 10
	now := time.Now()
	for i := 0; i < total; i++ {
		printout(i)
	}
	fmt.Println("\n", time.Now().Sub(now))
}
