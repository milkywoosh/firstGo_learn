package main
import "fmt"

//  can't be reassigned
const (
	data uint8 = 10
	data1 uint16 = 111
)
// var can be reassigned
var (
	dataVar string = "ganti"
)

func main() {


	fmt.Println(dataVar)

}