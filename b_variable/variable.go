package main

import "fmt"

//GLOBAL
const fixval = "value"

var val2 = 500

func main() {
	var age int = 100
	age2 := 200
	age3, age4 := 300, 400

	var name string
	name = "Gopher"
	var check bool

	fmt.Println(age, age2, age3, age4, name, check)
}
