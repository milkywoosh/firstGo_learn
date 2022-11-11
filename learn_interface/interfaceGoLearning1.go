package main

import(
	"fmt"
)
type Human interface {
	Eating(arg string) bool
}
type Muslim struct {
	name, religion string
}
type Kristen struct {
	name, religion string
}

func (m *Muslim) Eating(arg string) bool {
	var result bool
	if m.religion == "muslim" {
		if arg == "babi" || arg == "Babi" || arg == "BABI" {
			result = false
		} else {
			result = true
		}
	}
	return result
}

func (k *Kristen) Eating(arg string) bool {
	return true

}

func main() {
	// doni := &Muslim{"doni", "muslim"}
	doni := Muslim{"doni", "muslim"}
	robert := &Kristen{"robert", "kristen"}

	test_doni := doni.Eating("Babi")
	test_robert := robert.Eating("babi")
	fmt.Println(test_doni, test_robert)
}
