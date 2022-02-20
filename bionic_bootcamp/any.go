package main
import (
	"fmt"
	// other package
	// other package
)


type mod5Filter func(value int)int

func GetValue(value int, filter mod5Filter) {
	if filter(value) < 10 {
		fmt.Println("<10", value)
	} else {
		fmt.Println(">10", value)
	}
}

func filterValue(value int)int {
	return value
}

func main() {
	// append new element into arr from slc
	// var i int
	GetValue(11, filterValue)
	
	
}