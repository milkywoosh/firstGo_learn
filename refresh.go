// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"net/http"
)

type Animal struct {
	name string
}

func main() {
	/*
		f := 1
		f1 := Animal{}
		data, err := fmt.Fprintf(os.Stdout, "test %v %v\n", f, f1)
		if err != nil {
			fmt.Printf("error bos: %v\n", err)
		}
		fmt.Println(reflect.TypeOf(string(data)))
		fmt.Println(reflect.TypeOf(data))
	*/
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("error cuy %v\n", err)
	}
	// fmt.Printf("test cuy %v\n", resp)
	for _, data := range resp.Status {
		fmt.Printf("%v\n", data)

	}
}
