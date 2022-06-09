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

/*
	reference to read
	https://unitstep.net/blog/2015/09/16/golang-promoted-methods-method-sets-and-embedded-types/
*/
type Indonesia struct{}

func (i *Indonesia) Halo() {
	fmt.Println("Halo")
}

type Mexico struct{}

func (m *Mexico) Hola() {
	fmt.Println("Hola")
}

type England struct{}

func (e *England) Hello() {
	fmt.Println("Hello")
}

type language struct {
	Indonesia
	Mexico
	England
}

func main() {
	x := language{}
	x.Halo()
	x.Hola()
	x.Hello()
	x.England.Hello()

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
