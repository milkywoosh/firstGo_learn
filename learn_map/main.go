package main

import "fmt"

// goal -> create group of number type: int
type slc []int
type data_map map[string][]int
var store slc
var n int
var length int

func main() {
	var data data_map
	data = make(data_map)
	data["prime"] = []int{2,3,5,5,5,7,11,7,6,6, 5}
	data["odd"] = []int{1,3,5,7,9,11}
	data["even"] = []int{2,4,6,8,10}

	arr := data["prime"]
	for i, n := range arr {
		tmp:=i+1
		if arr[i] != arr[tmp] {
			
		}
	}

	fmt.Println(store)
}
