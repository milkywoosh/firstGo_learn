package main

import "fmt"

// accept multiple values
func SumN(arg ...int) int {
	result := 0
	for _, v := range arg {
		result += v
	}
	return result
}

func main() {
	// variadic function
	// spreading slice slice...
	data := []int{3, 3, 3, 3, 3}
	result := SumN(data...)
	fmt.Println(result)
}
