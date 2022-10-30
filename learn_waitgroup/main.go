package main

import (
	"fmt"
	"sync"
	_"testing"
	"time"
	_"net/http"
)


// using sleep is not IDEAL for running goroutine bcs it doesnt represent 'real condition'
/*
	running golang test
	--> go test
	--> go test -v
	--> go test -v -run TestNamaFunction
	
*/




func GoodExampleWG() {
	/*
	 unpredictable result
	 	-> total:  98
		   cek:  100
	*/

       
}

func main() {
	
	var wg sync.WaitGroup
	var total int = 0
	for i:=0;i<1000;i++ {
		wg.Add(1)
		go func(val int) {
			// fmt.Printf("%v\n", val)
			total+=1
			wg.Done()
		}(i)
		time.Sleep(1* time.Millisecond)
	}
	// wait all goroutine finish
	wg.Wait()
	fmt.Println("stop here !")
	fmt.Printf("%d\n", total)

        

}











