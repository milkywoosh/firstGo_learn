package main
import "fmt"

func Test(n uint16) uint16 {
	if n == 0 {
		return 0
	} else {
		return n + Test(n-1)
	}
}

func Looper(n uint16) uint16 {
	var data uint16 = n
	for {
		fmt.Println(data)
		data++
		if data == 5 {
			return data

		}
	}
}


func main() {

	fmt.Println("result: ", Looper(0))
	
}