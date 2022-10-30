package main

import(
	"fmt"
	"sync"
)

type Asset struct{
	mtx sync.Mutex
	money int
}
var wg sync.WaitGroup
// global variœable
var John Asset = Asset{
	money: 1000,
}
var Liz Asset = Asset{
	money: 100,
}

// init method print john's money using pointer
func (m *Asset) AddMoney(val int, wg *sync.WaitGroup) {

	
	m.mtx.Lock() // lock procces from mutex to avoid race condition
	m.money+=val
	defer m.mtx.Unlock() // unlock process from mutex to release go routine
	wg.Done()
}

func (m *Asset) CheckSaldo() {
	fmt.Printf("john's money: %v\n", m.money)
}


func main() {

	// use goroutine here
	for i:=0; i<1000000; i++ {
		wg.Add(1)
		go John.AddMoney(1, &wg)
	}

	wg.Wait()
	John.CheckSaldo()
}