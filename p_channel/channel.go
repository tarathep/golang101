package main

import (
	"fmt"
	"time"
)

func func1() {
	time.Sleep(1 * time.Second)
	fmt.Println("Success1")
}
func func2() {
	time.Sleep(1 * time.Second)
	fmt.Println("Success2")
}

func func3(message chan<- string) {
	time.Sleep(1 * time.Second)
	message <- "message" // insert value into Channel
}
func main() {
	start := time.Now()
	ch := make(chan string)
	go func1()
	go func2()
	go func3(ch) // ส่งท่อ ch เข้าไปใน Function (5)

	messgeFromFunc3 := <-ch           // ค่าจากท่อ Channel จะออกตรงนี้ (6)
	if messgeFromFunc3 == "message" { // (7)
		fmt.Println("Success3")
		fmt.Println(time.Since(start))
	}
}
