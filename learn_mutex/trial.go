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
	ch3 := make(chInt)
	flagger = 0

	go PassingChannel1(10, ch1)
	go PassingChannel1(100, ch2)
	go PassingChannel1(100, ch3)
	// time.Sleep(5 * time.Millisecond)


	for {
		select {
			// if ch1 read duluan
			case <- ch1:
				flagger+=1
				fmt.Println("ch1", ch1)
			// if ch2 read duluan
			case <- ch2:
				flagger +=1
				fmt.Println("ch2", ch2)

		}
		if flagger == 2 {
			break
		}
	}





}