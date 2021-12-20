package main
import (
	"fmt"
)

/*  
	create function here!! 
	dont forget to declare function type !!
*/

// cannot use "const"!!

var array = [3] string {"one", "two", "three"}

func main() {
	
	var i int
	for i=0; i<len(array);i++ {
		fmt.Println(array[i])
	}
	
}