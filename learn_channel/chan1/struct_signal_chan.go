package main

import(
	"fmt"
	"strings"
)

/*
	PURPOSE:
	- syncronization
	- using STRUCT TYPE for signaling channel, 
		syncronize to other goroutine
	
*/

func main() {
	done:= make(chan struct{})
	histogram:= make(map[string]int)
	data:= []string{
		"one two three", 
		"one two three",
		"one two three",
	}
	// fmt.Println("testing")
	go func() {
		defer close(done)
		for _, val:= range data{
			words:= strings.Split(val, " ")
			for _, word:= range words{
				// note: the map has been initialized, so it is started from 0
				// so it is automatically started from 0 and at the same time increment by 1
				histogram[word]+=1
			}
		}
	}()

	<-done // struct receiving to nothing only for signaling, and has no memory

	// for {
	// 	_, ok:= <-done // for syncronization from one goroutine to other goroutine
	// 	if !ok {
	// 		break
	// 	} else {
	// 		continue
	// 	}	
	// }
	fmt.Println("before unblocking")

	fmt.Println(histogram)
	




}













