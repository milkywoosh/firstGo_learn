package main

import "fmt"

type assets struct {
	vehicle []string
	liquid  []string
	land    []string
}
type cardholder struct {
	name    string
	deposit int
	assets  assets
}

type Cat struct{}

func (c *Cat) Meow() {
	fmt.Println("meow")
}

type Dog struct{}

func (d *Dog) Woof() {
	fmt.Println("woof")
}

type Wolf struct{}

func (w *Wolf) Auw() {
	fmt.Println("auww")
}

type animal struct {
	Cat
	Dog
	Wolf
}

func main() {

	x := animal{}
	x.Woof()
	x.Auw()
	x.Meow()

	/*
		var ben cardholder = cardholder{
			name:    "ben",
			deposit: 1000,
			assets: assets{
				vehicle: []string{"ferrari", "honda"},
				liquid:  []string{"gold", "bronze", "silver"},
				land:    []string{"california", "jakarta"},
			},
		}
		fmt.Printf("%+v\n", ben)
		fmt.Printf("%v\n", ben)
	*/

}
