package main

import (
		"fmt"
		_"time"
)


/*
	chan<- : send value to chan
	<-chan : receive from value

*/




// created func return send to channel [ chan<- ]

func main() {

	words := []string{"foo", "bar", "baz"}
	done := make(chan string, len(words))
	for _, word := range words {
	    go func(word string) {
	       // a job
	        fmt.Println(word)
	        done <- word// not blocking
	    }(word)
	}
	
	for _, i := range words {
	    fmt.Println(<-done, i)
	}


}