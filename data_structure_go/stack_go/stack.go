package main

import (
	"fmt"
)

type lifo_slice []string

func main() {
	data := lifo_slice{"one", "two", "three"}
	n := len(data) - 1
	for n >= 0 {
		fmt.Println(data[n])
		n -= 1
	}
}
