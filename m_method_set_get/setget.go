package main

import "fmt"

//method on primitive type

type text string

func (t *text) set(s string) {
	*t = text(s)
}

func (t *text) get() string {
	return string(*t)
}

func main() {
	var t text
	t.set("sss")
	fmt.Println(t.get())

}
