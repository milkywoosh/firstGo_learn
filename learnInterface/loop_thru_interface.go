// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Predator interface {
	Hunting() string
}
type Mamals struct {
	name string
}
type Insect struct {
	name string
}

func (m *Mamals) Hunting() string {
	return "hunting....."
}
func (i *Insect) Hunting() string {
	return "hunting....."
}

func main() {
	// need ampersand -> "&" if each method has "*" pointer ln17 ln20
	croco := &Mamals{"croco"}
	tiger := &Mamals{"tiger"}

	data := []Predator{croco, tiger}
	for _, do := range data {
		fmt.Println(do.Hunting())
	}
}
