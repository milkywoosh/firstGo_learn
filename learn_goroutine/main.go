package main

// `go test -v -run=TestFunction`

import (
	"fmt"
	"time"
)

func cha (v int, ch chan int) {
	time.Sleep(100 * time.Millisecond)
	ch <- v
}

func chb (v int, ch chan int) {
	time.Sleep(100 * time.Millisecond)
	ch <- v
}

var chfinal int
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go cha(100, ch1)
	go chb(100, ch2)
	
	chfinal = <-ch1 + <-ch2
	fmt.Println(chfinal)
	



}