package main

import(
	"fmt"
)

type Bot struct {
	name string
	type_bot string
}
type MultipleBot struct {
	data []*Bot
}

func main() {
	bot1 := &Bot{
		name: "bot1",
		type_bot: "human",
	}

	fmt.Println(bot1)
}