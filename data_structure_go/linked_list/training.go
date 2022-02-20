package main
import "fmt"
type Node struct {
	prev *Node
	next *Node
	key  int
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key int) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		// fmt.Println("list-> ", *list)
		L.head.prev = list
	}
	L.head = list

	currNode := L.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	L.tail = currNode
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}


func main() {
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)
	
	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	fmt.Println("\n==============================\n")
/*
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	link.Reverse()
	fmt.Println("\n==============================\n")
*/
}