package main
import "fmt"

const (
	firstState bool = true
)

/*
func BinarySearch(n int, start int, end int, array[] int) [] int {

}
*/

func GetArray(arr[3] int) [3]int {
	return arr
}

func main() {
	// var test = [...]int{1,2,3}
	fmt.Println(GetArray([3]int{4,3,2}))

}