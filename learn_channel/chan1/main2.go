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
	bot2 := &Bot{
		name: "bot2",
		type_bot: "human",
	}

	group_of_bot := &MultipleBot{
		data: []*Bot{
			bot1,
			bot2,
		},
	}

	for _, val := range group_of_bot.data {
		fmt.Println(val)
	}
}