package main
import "fmt"


const (
	salary int = 1000
	tax float32 = 0.05
)

func main() {
	salary:= float32(salary)
	var result float32 = salary*tax
	fmt.Println(result) 

}