package main

import(
	fmt "fmt"
	time "time"
)

func main() {
	

	for i:=0; i<100; i++ {

		go func() {
			fmt.Println("1", i)
		}()
		go func() {
			fmt.Println("2", i)
		}()
	}

	time.Sleep(5 * time.Millisecond)
}