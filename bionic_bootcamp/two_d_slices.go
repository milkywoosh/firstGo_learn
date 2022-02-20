package main
import (
	"fmt"
)

/*
	{
		{x,x,x,x}, <== 2 rows, 4 columns
		{y,y,y,y}
	}
*/

//2D slices
var two_d_slices = [][] int { {1,2,3}, {4,5,6} }

var rows int = 3
var cols int = 5
var tes2d = make([][]int, rows)

var arr = []int{1,2,3,4}
var add = []int{7,7,7}

func main() {
	// fmt.Println(tes2d)
	var i int
	/*
	for i = range(tes2d) {
		if (i == 1) {
			tes2d[i] = []int{1,2,3}
		} else {
			tes2d[i] = []int{0,0,0}
		}
	}
	*/
	for i = range(tes2d) {
		tes2d[i] = make( []int, cols )
	}

	// slicing element
	// fmt.Println(arr[:5])

	// --> add many element from other slice
	for i = range add {
		arr = append(arr, add[i])
	}

	for i = range arr {
		
	}
}




