package main


/*
	Given a nil channel, this is several case to understand
	note: 
	>  <-c [receiving from c] : blocks forever
	>  c<- v [sending to c] : blocks forever
	> close(c) [close nil c] : [panic]

	now we'd like to handle this condition
	reference: 
		==> https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308
*/

import (
	"fmt"
	t "time"
	_"math/rand"
)

// accept multiple int values
func asChan (vs ...int) <-chan int {
	// init non nil channel
	c := make(chan int) // bidirectional channel
	go func() {
		for _, v := range vs{
			c <- v
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			t.Sleep(1000 * t.Millisecond)
		}
		close(c)
	}()
	return c
}
func merge(a, b <-chan int) chan <- int {
	c := make(chan int)

	go func() {
		for {
			select {
			// v is variabel that receive from channel 
			case v:= <- a:
				c <- v
			case v:= <- b:
				c <- v
			}

		}
	}()
	return c
}

func main() {
	a:= asChan(1,3,5,7)
	b:= asChan(2,4,6,8)
	c:= merge(a,b)
	
	for v:= range c {
		fmt.Println(v)
	}

}