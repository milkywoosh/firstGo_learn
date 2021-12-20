package main
import (
	"fmt"
)

// declare array
var Employee = [3] string{"Ben", "Ron", "John"}

// declare function
func TesSwitch(arg[3] string, target int) string {
	var result string
	for i:=0; i<len(arg); i++ {
		if (i == target) {
			switch (arg[target]) {
				case "Ben":
					result = "ONE"
					break
				case "Ron":
					result = "TWO"
					break
				case "John":
					result = "THREE"
					break
				default:
					result = "NOTHING"
					break
			}
		} else {
			continue
		}
		
	}
	return result
}

func main() {
	
	result:=TesSwitch(Employee, 2)
	fmt.Println(result)
}

