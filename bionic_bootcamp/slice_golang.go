package main
import (
	"fmt"
)
/*
	[] type_value
	append(arr, value)
	len()
	cap()
*/

var data = [] int {1,2,3,4}
// modifier function
func twiceValue(slice [] int) {
	var index, value int
	for index, value = range(slice) {
		slice[index] = value*10
	}
}

func main() {
	twiceValue(data)
	fmt.Println(data)
}