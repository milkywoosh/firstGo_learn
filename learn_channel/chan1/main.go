package main

import (
	"fmt"
	"time"
)

// create channel receiver with function -> done
// create channel sender with function -> done
// create channel buffer (chan string, 3) -> done
	// note: buffer in goland should reflect maximum channel should be received or sent
	// if using goroutine, it will give zero to amount result if printed
// can i use like this --> (chan [custom type]


func Test(val int, ch chan int) {
	ch <-val
	ch <-val
	ch <-val

	time.Sleep(5 * time.Millisecond)
	defer close(ch)
}


func main() {
//	fmt.Println("hello world")
	ch1 := make(chan int)
//	go Test(5, ch1)
	defer close(ch1)



	ch1 <- 10
	fmt.Println("test")


}

