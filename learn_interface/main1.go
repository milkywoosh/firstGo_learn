package main

import(
	"fmt"
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
	Bot1.Add_skill("archery")
	Bot1.Add_skill("swimming")
	Bot1.Add_skill("debating")
	fmt.Println(Bot1)
	fmt.Println(Bot1.Get_all_skill())

}













