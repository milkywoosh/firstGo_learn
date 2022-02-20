package main
import (
	"fmt"
)
/*
	note: Go can automatically increase its "capacity"
	from 2,4,8,16,32 and so on..
	after the curent length increase more than
	current capacity
*/

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