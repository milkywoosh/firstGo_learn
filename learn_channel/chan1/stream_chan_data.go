package main
import(
	"fmt"
	"strings"
)

/*

	streaming DATA with channel
	- data from MAIN GOROUTINE	
	- data goes to SUBGOROUTINE
	- data goes back to MAIN GOROUTINE
	- data in MAIN GOROUTINE is ready

*/


func main() {
	ch:= make(chan string)
	data:= []string{
		"one two three", 
		"one two three",
		"one two three",
	}
	histogram:= make(map[string]int)
	/*

	*/
	go func() {
		defer close(ch) // question, when is actually CHANNEL CLOSE?????
		for _, value:= range data{
			words:= strings.Split(value," ")
			for _, word:= range words {
				// histogram[word]+=1 
				// try to replace supply each word with channel
				ch<- word
				

			}
		}
	}()
	// easiest way --> the cleanest way to receive from channel
	/*
	for val:= range ch {
		fmt.Println(val)
		histogram[val]+=1
	}
	fmt.Println(histogram)
	*/
	

	// second easiest way using 'break' [val, ok:= <-ch]
	/*
		for {
			val, ok:= <-ch
			if !ok {
				fmt.Printf("%v\n", "stop -------")
				break 
			} else {
				histogram[val]+=1
			}
		}

		fmt.Println(histogram)
	*/






}