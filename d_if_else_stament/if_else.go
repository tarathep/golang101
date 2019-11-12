package main

import "fmt"

func main() {
	condition := true

	if condition {
		fmt.Print("A")
	}
	if condition == true {
		fmt.Print("B")
	} else {
		fmt.Println("C")
	}

	if condition {
		if condition != true {
			fmt.Print("D")
		} else {
			fmt.Print("E")
			if v := false; v == condition {
				fmt.Print("F")
			}
		}
	}

	val := 200
	switch val {
	case 100:
		fmt.Print("case1")
		break //
	case 200:
		fmt.Print("case2")
	case 300:
		fmt.Print("case3")
	default:
		fmt.Print("default")
	}
}
