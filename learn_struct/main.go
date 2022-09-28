package main

import (
	"encoding/json"
	"fmt"
)

type any map[string]int
type habitat_detail struct {
	Continent []string `json:"continent"`
	Land_type string   `json:"land_type"` // aqua, forest, desert, grassland, tundra
}
type animal struct {
	Weight  int8           `json:"weight"`
	Height  int8           `json:"height"`
	Class   string         `json:"class"`
	Habitat habitat_detail `json:"habitat"`
	Anydata any
}

// declare here

var cat animal
var cat_habitat habitat_detail
var any_data any

func main() {

	any_data = make(any)
	any_data["one"] = 1
	any_data["two"] = 2

	cat_habitat = habitat_detail{
		Continent: []string{"asia", "europe", "america", "africa"},
		Land_type: "grassland",
	}

	cat = animal{
		Weight:  10,
		Height:  5,
		Class:   "mamals",
		Habitat: cat_habitat,
		Anydata: any_data,
	}

	// convert to json, and see with Marshalindent
	datajson, err := json.MarshalIndent(cat, "_", " ")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("%+v\n", string(datajson))

}
