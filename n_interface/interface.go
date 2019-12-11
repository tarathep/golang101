package main

import "fmt"

//รูปร่าง
type Shape interface {
	Area() float64      //พื้นที่
	Perimeter() float64 //เส้นรอบรูป
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	//var s Shape
	//s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	//fmt.Printf("type of s is %T\n", s)
	//fmt.Printf("value of s is %v\n", s)
	//fmt.Println("area of rectange s", s.Area())
	fmt.Println(r.Area(), r.Perimeter())
	fmt.Println(Rect{5.0, 4.0}.Area())
	//fmt.Println("s == r is", s == r)
}
