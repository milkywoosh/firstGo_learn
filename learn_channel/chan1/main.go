package main



// create channel receiver with function -> done
// create channel sender with function -> done
// create channel buffer (chan string, 3) -> done
	// note: buffer in goland should reflect maximum channel should be received or sent
	// if using goroutine, it will give zero to amount result if printed
// can i use like this --> (chan [custom type]


// ---> DONT DELETE
// 	   ch:= make(chan int)// [bidirectional channel] if this is NOT BUFFERED -> must used inside new goroutine
//     ch1:= make(chan int)
//     // 1 of way to handle proper channel
//     go func() {
//     	ch <- 10
//     }()

//     v:= <-ch
//     fmt.Println("v: ", v)
//     go func(val int) {
//     	ch1 <- val
//     }(v)
    
//     v1:= <-ch1
//     fmt.Printf("v1: %v\n", v1)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

// func Foo(ch chan int, v ...int) {	
// 	for _, i := range v {
// 		ch <-i
// 	}
// 	close(ch)
// }

// func main() {
// 	ch:= make(chan int, 5)
	
// 	go Foo(ch,2,3,4,5)
	
// 	for i:= range ch{
// 		fmt.Println(i)
// 	}
// }


// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	// ch:= make(chan int)
	// data:= []int{1,2,3,4,5,6}
	// go func() {
	// 	defer close(ch)
	// 	for _, val := range data {
	// 		ch<-val
	// 	}
	// }()

	// for i:= range ch {
	// 	fmt.Println(i)
	// }
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// func isClosed(ch <-chan int)bool {
// 	select {
// 	case v:=<- ch:
// 		fmt.Println("test", v)
// 		return true
// 	default :
// 		fmt.Println("test false")
// 		return false
// 	}
// }

// func main() {
	
// 	ch:= make(chan int)
// 	fmt.Println(isClosed(ch))
// 	close(ch)
// 	fmt.Println(isClosed(ch))

	
	
// }
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
import (
	"fmt"
	"time"
	"math/rand"
	_"log"
)


func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					fmt.Println("a is done")
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					fmt.Println("b is done")
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}


func merge_failed(a, b <-chan int) <-chan int {
c := make(chan int)
	go func() {
        defer close(c)
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
                    fmt.Println("a is done")
					adone = true
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
                    fmt.Println("b is done")
					bdone = true
					continue
				}
				c <- v
			}
		}
	}()
	return c
}
func main() {
	
a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
	
	
}



