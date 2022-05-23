package main

import (
	"fmt"
)

// first in first out ==> loop biasa
// first in first out ---> reversed loop
type sliceData []string
type hashMap map[string]sliceData

func main() {

	data := sliceData{"one", "two", "three"}

	test := hashMap{}
	test["one"] = data
	test["two"] = data

	fmt.Println(test)

}
