package main

import(
	"fmt"
)

type Human struct {
	name string
	age, phone int
}

type Student struct {
	Human
	school string
	loan int32
}

type Employee struct {
	Human
	company string
	money float32
}

// when variable is Men type, it MUST implement 
// those 2 method !
type Men interface {
	SayHi()
	Sing(lyrics string)
}

// method struct for implementing interface
// must NOT use receiver type
// must have/implement all struct methods declared on interface
func (s Student) SayHi() {
	fmt.Printf("Hello from %s !\n", s.name)
}
func (s Student) Sing(lyrics string) {
	fmt.Printf("%s like to sing %s\n", s.name, lyrics)
}

func (h Human) Talk() {
	fmt.Printf("Hello my name is %s\n", h.name)
}


var ben Men // Men is type interface, must 
var penny Men // Error because hasnt have method declared on Men interface
func main() {
	ben = Student{Human{"ben", 19, 10029},"math", 100}
	// ben.SayHi()
	// ben.Sing("Oasis song")


	penny = Human{"penny", 14, 2345}
	penny.SayHi()

}