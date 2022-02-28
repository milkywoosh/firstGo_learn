package main

import (
	"fmt"
)

// reference to read --> https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b

func ModifByReference(value *float64) (data float64) {
	data = *value * 2
	return
}

func main() {
	x := 1.3
	result := ModifByReference(&x)
	fmt.Println(result)
}
