package main

import(
	"fmt"
	"time"
)

type Bot struct {
	name string
	type_bot string
	skills []string
}

func (b Bot) Get_all_skill() string{
	store := ""
	for i, val := range b.skills {
		if i == len(b.skills)-1 {
			store+= val
		} else {
			store+= val+", "	
		}
		
	}
	return fmt.Sprintf("%v has the following skills %v\n", b.name, store)
}
func (b *Bot) Add_skill(val string) {
	b.skills = append(b.skills, val)
}


func main() {

	var Bot1 Bot
	Bot1 = Bot{
		name: "Bot1",
		type_bot: "Servant",
		skills: []string{"cooking"},
	}
	go Bot1.Add_skill("archery")
	go Bot1.Add_skill("swimming")
	go Bot1.Add_skill("debating")

	time.Sleep(500 * time.Millisecond)
	fmt.Println(Bot1)
	fmt.Println(Bot1.Get_all_skill())
}







