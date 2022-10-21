package main

import (
	"fmt"
	"sync"
	_"testing"
	_"time"
	_"net/http"
)


// using sleep is not IDEAL for running goroutine bcs it doesnt represent 'real condition'
/*
	running golang test
	--> go test
	--> go test -v
	--> go test -v -run TestNamaFunction
	
*/




func GoodExampleWG_2() {
        var wg sync.WaitGroup
        var total = 0
		var cek = 0

		caller_total := func (n int, val_ref *int, wg *sync.WaitGroup) {
			   defer wg.Done()
                        *val_ref+=1
		}
		caller_cek := func (n int, val_ref *int, wg *sync.WaitGroup) {
			   defer wg.Done()
                        *val_ref+=1
		}

        for i:=0; i<100; i++ {
                // Increment the WaitGroup counter.
                wg.Add(2)
                // Launch a goroutine to fetch the URL.
              	go caller_total(i, &total, &wg)
              	go caller_cek(i, &cek, &wg)
        }
        // Wait for all HTTP fetches to complete.
        wg.Wait()
        fmt.Println("total: ", total)
        fmt.Println("cek: ", cek)
}

func main() {
	for i:=0; i<10; i++ {
		GoodExampleWG()
	}
}











