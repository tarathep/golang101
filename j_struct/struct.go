package main

import (
	"fmt"
)

type Book struct {
	title    string
	author   string
	subtitle string
	price    float64
}

func main() {
	var golang Book
	golang.title = "Go Programming"
	golang.author = "tarathep"
	golang.subtitle = "go tutorial"
	golang.price = 199.99
	fmt.Println(golang)

	golang2 := Book{
		title:    "Go Programming",
		author:   "tarathep",
		subtitle: "go tutorial",
		price:    200.00,
	}

	golang3 := Book{"Go", "bokie", "tutorial", 299.0}

	fmt.Println(golang2.price * 2)
	fmt.Println(golang3)

	//mainx()
}

/*
func mainx() {
	jsonBlob := []byte(`{
        "width": 48,
        "length": 64,
        "border": "solid"
}`)

	type rectangle struct {
		Width   int    `json:"width"`
		Length  int    `json:"length"`
		Borders string `json:"border"`
	}

	var rec rectangle
	err := json.Unmarshal(jsonBlob, &rec)
	if err != nil {
		log.Fatal(err)
	}
}
*/
