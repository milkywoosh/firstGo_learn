package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head   *Node
	length int
}

// method struct
func (l *LinkedList) AddToTail(value int) {
	n := Node{}
	n.data = value
	//fmt.Println(n)
	if l.length == 0 {

		l.head = &n
		fmt.Println(*&l.head)
		l.length++
	} else {
		link := l.head
		for link.next != nil {
			link = link.next
		}
		link.next = &n
		fmt.Println("---> ", *&link.next)
		l.length++
	}
}

//void function
func (l *LinkedList) Display() {
	if l.length == 0 {
		fmt.Println("data is empty")
	} else {
		link := l.head
		fmt.Println(link.data)
		for link.next != nil {
			link = link.next
			fmt.Println(link.data)
		}
	}
}

// func main() {
// 	listData := LinkedList{}
// 	listData.Display()
// 	listData.AddToTail(10)
// 	listData.AddToTail(12)
// 	listData.AddToTail(14)
// 	listData.Display()
// }
