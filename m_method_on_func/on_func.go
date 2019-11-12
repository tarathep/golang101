package main

import "fmt"

/*
type text func() string

func (t text) get() string {
	return t()
}

func lambda() string {
	return "message function"
}

func main() {
	//    var t text = lambda
	t := text(lambda)
	fmt.Println(t.get())

}
*/

type text func() string

func (t text) get() string {
	return t()
}

func main() {
	var t text = func() string {
		return "message function"
	}

	fmt.Println(t.get())

}
