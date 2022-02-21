package main

import "fmt"

// accept multiple values argument
/*
input slice[] int
each n converted to float64
return value float64

*/
func AverageN(arg ...int) (result float64) {
	result = 0.0
	for _, v := range arg {
		result += float64(v) // convert to float64
	}
	result /= float64(len(arg))
	return
}

func main() {
	// variadic function
	// spreading slice slice...
	data := []int{2, 3, 6, 1, 7}
	result := AverageN(data...)
	fmt.Println(result)

}
