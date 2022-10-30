// create select channel after being the receiver
package main

import (
	"fmt"
)

type chint chan int
type Wallet struct {
	v int
	user string
}

// add money to ben using channel
var Ben Wallet = Wallet {
	v : 1000,
	user : "ben",
}

var Add chint = make(chInt)
var Subtract chint = make(chInt)

// struct method for Ben
func (w *Wallet) AddMoney(v int, sent chint) {

	for {
		select {
		case <- 
		}
	}
}

func main() {
	fmt.Printf("%v\n", Ben)


}