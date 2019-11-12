package main

import "fmt"

func main() {
	//MAP (DICT) find by string key get val
	n := make(map[string]string)
	n["TH"] = "THAILAND"
	n["JP"] = "JAPAN"
	n["EN"] = "ENGLAND"

	fmt.Println(n["JP"])

	y := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
		"B":  "Boron",
	}

	fmt.Println(y["H"])
}
