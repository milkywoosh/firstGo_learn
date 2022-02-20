package main
import (
	"fmt"
)

type FilterData func(string) string

func getData(data string, filter FilterData) string {
	result := filter(data)
	return result
}

// void --> without any return value
func filterDataHelper(data string) string{
	if data == "Anjing" || data == "Babi" {
		return "..."
	} else {
		return data
	}
}

func main() {
	// result := getData("Anjing", filterDataHelper)
	result := getData("Babi", filterDataHelper)
	fmt.Println(result)
}