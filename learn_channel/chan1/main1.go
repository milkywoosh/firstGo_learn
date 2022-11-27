package main

import (

	"fmt"
	"time"
)

var StoreReceiver int

func PassChan(v int, chEven chan int, chOdd chan int) {

	if v%2 == 0 {
		fmt.Println("even")
		// time.Sleep(2 * time.Second)
		chEven <-v
		
		//defer close(chEven)
	} else {
		fmt.Println("odd")
		// time.Sleep(2 * time.Second)
		chOdd <- v

		//defer close(chOdd)
	}

	defer close(chEven)
	defer close(chOdd)
}



func main() {
	//

	chEven_ := make(chan int,1)
	chOdd_ := make(chan int,1)

	go PassChan(18, chEven_, chOdd_)
	time.Sleep(1 * time.Second) // --> give chance to variable receiver to receive value

		select {
			case StoreReceiver= <- chEven_:
				 fmt.Println("chEven_")
				 fmt.Println(StoreReceiver)

			case StoreReceiver=<- chOdd_:
				fmt.Println("chOdd_")
				fmt.Println(StoreReceiver)
			default :
				fmt.Println("no data received??")
		}

//	defer close(chEven)
//	defer close(chOdd)

}
