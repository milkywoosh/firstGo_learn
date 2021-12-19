package main
import "fmt"

const (
	data1 uint16 = 100
	data2 uint16 = 200 
)

func main() {

	switch (true) {
		case (data1 > data2):
			fmt.Println("data1 > data2")
		case (data1 < data2):
			fmt.Println("data1 < data2")
		case (data1 == data2):
			fmt.Println("data1 == data2")
		default:
			/* code */
			fmt.Println("nothing, it's default")
	}
}