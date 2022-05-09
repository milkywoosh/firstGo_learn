// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	f := 1
	f1 := 2
	data, err := fmt.Fprintf(os.Stdout, "test %v %v", f, f1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println(reflect.TypeOf(data))

}
