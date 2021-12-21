package main
import (
	"fmt"
)


var arr = [3] int {1,2,3}
func main() {

	/*	var j int
		for j=0; j<5; j++ {
			fmt.Println(j)
		}
	*/

	var i, value int
	for i, value = range (arr){
		fmt.Printf("Index: %d -> value: %d\n", i, value)
	}

}