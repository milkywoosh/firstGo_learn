package main

import (
	fmt "fmt"
	time "time"
)

type chInt chan int

var val int
var temp func(int)int
var flagger int

func PassingChannel1(val int, ch chInt) {
	time.Sleep(5 * time.Millisecond)
	ch <- val

}


func main() {
	ch1 := make(chInt)
	ch2 := make(chInt)
	flagger = 0

	go PassingChannel1(10, ch1)
	go PassingChannel1(100, ch2)
	// time.Sleep(5 * time.Millisecond)


	for {
		select {
			case <- ch1:
				flagger+=1
				fmt.Println("ch1")

			case <- ch2:
				flagger +=1
				fmt.Println("ch2")

		}
		if flagger == 2 {
			break
		}
	}





}