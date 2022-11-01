package main

import (
	"fmt"
	"sync"
)

func main() {

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// hello := func(wg *sync.WaitGroup, id int) { 
	// 	defer wg.Done()
 //        fmt.Printf("Hello from %v!\n", id)
 //    }
	// const numGreeters = 20
	// var wg sync.WaitGroup 
	// wg.Add(numGreeters)

	// for i := 0; i < numGreeters; i++ {
	// 	go hello(&wg, i+1) 
	// }
	// wg.Wait()
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++


	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	var count int
	var lock sync.Mutex
	increment := func() { 
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() { 
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}
	// Increment
	var arithmetic sync.WaitGroup 
	for i:=0;i<=5;i++{
		arithmetic.Add(1) 
		go func() {
			defer arithmetic.Done()
	        increment()

	    }()
	    arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
	        decrement()
     	}()
	}
	
	
    arithmetic.Wait()
	fmt.Println("Arithmetic complete.")

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
}