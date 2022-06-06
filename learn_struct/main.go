// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

type machine struct {
	name          string
	energy_source map[string]string
}

// modify machine energy source
/*
func (m *machine) Modify(key string, newValue string) {
	if energy_source[key] == nil {
		fmt.Println("data not found")
	}
	m.energy_source[key] = newValue
}
*/

func main() {
	var data *machine
	data = &machine{}
	data.name = "ben"
	data.energy_source = make(map[string]string) // need to create instance of map first !!
	//	data.energy_source["one"] = "satu"
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data))

	var data1 machine = machine{}
	fmt.Println(data1)
	fmt.Println(reflect.TypeOf(data1))

}
