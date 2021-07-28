package main
import "fmt"


/*
	data type:
		int
		float32
		string
		bool
*/

/*
	Printf() Function
	%v  --> to get value
	%T --> to get type

*/

var animal, human, tree, car string = "cat", "john", "oak", "tesla"

func main() {

	var name string = "ben"
	var age int = 14
	var isFriendly bool = !(true && true)

	name = "ron" // reassign name value

	var empty_string string
	var empty_int int
	var empty_bool bool
	var empty_float float32

	// one time "var"
	var (
		satu int = 1
		dua int = 2
		tiga int = 3
	)

	idk := 10


	fmt.Println(age)
	fmt.Println(isFriendly)
	fmt.Println(name)
	fmt.Println("my words!")

	fmt.Println(empty_string)
	fmt.Println(empty_int)
	fmt.Println(empty_bool)
	fmt.Println(empty_float)

	fmt.Println(idk)

	fmt.Println(animal)
	fmt.Println(human)
	fmt.Println(tree)
	fmt.Println(car)

	fmt.Println(satu, dua, tiga)

	// Printf
	fmt.Printf("this name: %v and this type: %T \n", name, name)




	// input process
	var getInput int
	fmt.Print("type number input: ")
	fmt.Scan(&getInput)
	fmt.Println("result input: ", getInput)

}