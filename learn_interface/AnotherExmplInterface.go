// You can edit this code!
// Click here and start typing.
package main

import "fmt"


// ======> nice thread : https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

// interface for method contract??
// group of methods
type Calculate interface {
	NetIncome() int
	TaxBill(persen float64) float64
}
type Amartha struct {
	revenue, cost int
}
type Credivo struct {
	revenue, cost int
}

func (a Amartha) NetIncome() int {
	return a.revenue - a.cost
}
// function with casting int to float64
func (a Amartha) TaxBill(persen float64) float64 {
	return float64(a.revenue) * persen / 100.0
}
func (c Credivo) NetIncome() int {
	return c.revenue - c.cost
}
func (c Credivo) TaxBill(persen float64) float64 {
	return float64(c.revenue) * persen / 100.0
}

// arg for interface
func ExecInterface(c Calculate) {
	fmt.Println("check: ", c)
	fmt.Println("income: ", c.NetIncome())
	fmt.Println("tax: ", c.TaxBill(2))
	//fmt.Println()
	//fmt.Println()
}
func main() {
	// have contract interface with Calculate
	// revenue int, cost int
	Am := Amartha{1000, 200}
	Cr := Credivo{400, 600}
	// implement interface with all methods in Calculate
	ExecInterface(Am)
	ExecInterface(Cr)
}
