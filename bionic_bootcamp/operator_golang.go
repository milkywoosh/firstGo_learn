package main
import "fmt"

const (
	x uint8 = 8
	y uint8 = 4
)

func main() {
	fmt.Printf("Penjumlahan %d\n", x + y)
	fmt.Printf("Pengurangan %d\n", x - y)
	fmt.Printf("Perkalian %d\n", x * y)
	fmt.Printf("Pembagian %d\n", x / y)
}