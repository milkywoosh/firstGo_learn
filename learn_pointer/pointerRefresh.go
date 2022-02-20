// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var data int = 5 // stored in random access memory
var dataX int = 10

func main() {
	// learning pointer
	/*
		test := &data // get memory of data variable
		fmt.Println(test)
		fmt.Println(*test) // get value stored in memory
		*test = *test * 3  // modify without copy value
		fmt.Println(*test)
	*/
	fmt.Println("\ncopy value to new variable, LESS EFFECTIFE IN MEMORY USAGE\n")
	// copy value to new variable, LESS EFFECTIFE IN MEMORY USAGE
	test1 := data
	for test1 < 10 {
		test1++
	}
	fmt.Println("this is test1: ", test1)
	// now print data 7th line ---> presumably NOT CHANGE bcs test1 COPIED value and NOT REFERENCING !
	fmt.Println("\nthis is data: ", data, "as the ancestor value")
}
