package main
import (
	"fmt"
	_ "reflect"
)


// declare type 
type slc_str []string
type data_map map[string]slc_str
type base_map map[string]int
type nested_map map[string]base_map

var test_map nested_map
// var tes_data_map nested_map
var consumer_type = 
func Looper(data data_map) {
	for i, val := range data["mamals"] {
		fmt.Println(i, val)
	}
}
func main() {
	// review golang
	test_map = nested_map{}
	test_map["asset"] = base_map{}
	test_map["asset"]["gold"] = 10
	test_map["asset"]["land"] = 2000

	for _, val := range test_map {
		for _, sub_val := range val {
			fmt.Println(sub_val)
		}
	}
	



}