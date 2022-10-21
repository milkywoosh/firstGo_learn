package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


// using sleep is not IDEAL for running goroutine bcs it doesnt represent 'real condition'
/*
	running golang test
	--> go test
	--> go test -v
	--> go test -v -run TestNamaFunction
	
*/


// to run
func Woke(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1)
	
}

var total = 0
func main() {


	for i=0; i<100; i++ {
		go func() {
			total+=1
		}()
	}


	fmt.Printf("%v\n", total)

}













