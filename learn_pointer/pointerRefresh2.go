// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

var name string = "john"
var num int = 5
var num1 int = 10

var a *int // 0x5071c8 => &num

func Modify(b *int) {
	fmt.Println(b)
}

func GetAddressBaseString(arg *string) {
	fmt.Println(arg)
}

func Square(v int) {
	v = v * v
	fmt.Printf("test %v %v", &v, v)
}

func main() {
	// need arg type pointer base int
	p0 := new(int)

	x := p0
	fmt.Println(x)

	fmt.Println(**&x)

	fmt.Println(reflect.TypeOf(x))
}
