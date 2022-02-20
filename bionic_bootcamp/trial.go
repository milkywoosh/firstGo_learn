package main
import (
	"fmt"
)

// slice 
/*

	append()
	len()
	cap()
*/

var arr = [4]int{2,2,2,2}
var addition =  []int{1,1,1,1,1,1,1,1}
func main() {
	var i int

	slice:= arr[0:2]
	fmt.Println(slice)
	fmt.Println(cap(arr))
	fmt.Println("cap(slice): ", cap(slice))	
	slice = append(slice, 4)
	slice = append(slice, 10)
	slice = append(slice, 20)
	fmt.Println("cap(slice): ", cap(slice))	
	

	// increasing capacity to 32?? by adding 8 more element, 
	// from existing 9 element
	for i = range addition {
		slice = append(slice, addition[i])
	}
	fmt.Println("len: ",len(slice))
	fmt.Println("cap(slice): ", cap(slice))	
	fmt.Println(slice)

}








