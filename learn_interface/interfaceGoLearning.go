// need to copy interfaceBasic.go
package main

import (
	"fmt"
)

type Standardized interface {
	FilterValue(standardValue int) (bool, string)
	MultiplyLimit(standardValue int, consta int) bool
}

type Sembako struct {
	name   string
	weight int
}

// struct method
func (s Sembako) FilterValue(standardValue int) (bool, string) {
	if s.weight > standardValue {
		return false, "too much"
	} else {
		return true, "fair"
	}

}
func (s Sembako) MultiplyLimit(standardValue int, consta int) bool {
	if s.weight*2 > standardValue {
		return false
	} else {
		return true
	}
}

func ResumeResult(s Standardized) {
	fmt.Println(s)
	a, b := s.FilterValue(70)
	fmt.Println(a, b)
}

func main() {
	beras := Sembako{"beras", 50}
	//c := beras.MultiplyLimit(120, 2)
	ResumeResult(beras)
}
