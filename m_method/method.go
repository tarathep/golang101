package main

import "fmt"

type rectangle struct {
	width  int
	length int
}

func (rec rectangle) area() int {
	return rec.width * rec.length
}

/*func area(rec rectangle) int {
	return rec.width * rec.length
}
*/
func main() {
	//rec := rectangle{3, 4}
	//fmt.Println(area(rec))
	rec := rectangle{3, 4}
	fmt.Println(rec.area())
}
