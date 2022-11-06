package main

import(
	"fmt"
)

// question: WHY method interface cannot use pointer type?? reference -> https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver

// SOURCE LEARNING: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/
// book build-web-application-with-golang-en !! in learn_golang folder

type Transportation struct {
	passanger string
	track string
	load int
	fuel string
}
type Car struct {
	Transportation
	wheel int
	owner string
	amount_customer int
}
type Motorbike struct {
	Transportation
	owner string
	automatic bool
}
// when variable is CityTrans type, it MUST implement \n
// those 2 method !
type CityTrans interface {
	Courrier(goods string)
	RideHailing(customers string, destination string)
	TotalCustomer()int
}
// method struct for implementing interface
// must NOT use receiver type
// must have/implement all struct methods declared on interface
func (c *Car) Courrier(goods string) {
	fmt.Printf("%s currently is delivering %s\n", c.owner, goods)
}
func (c *Car) RideHailing(person string, destination string) {
	c.amount_customer+=1
	fmt.Printf("%s currently is picking %s to %s\n", c.owner, person, destination)
}
func (c *Car) TotalCustomer()int{
	return c.amount_customer
}


// var ben Men // Men is type interface, must 
// var penny Men // Error because hasnt have method declared on Men interface
func main() {
	
	var car1 CityTrans
	car1_ := Car{
		Transportation: Transportation{
			passanger: "worker",
			track: "land",
			fuel: "pertalite",
		},
		wheel: 4,
		owner: "car1",
		amount_customer: 0,
	}
	car1 = &car1_

	// car1.Courrier("food")
	car1.RideHailing("gen", "lubang buaya")
	car1.RideHailing("ron", "PIK")
	car1.RideHailing("ken", "kokas")
	fmt.Printf("currently %s has reached %v customers\n", car1_.owner , car1.TotalCustomer())

	
}
















