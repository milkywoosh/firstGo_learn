// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// pointer
type Node struct {
	data int
	next *Node
}
type Link struct {
	head *Node
	len  int
}

// struct method
func (l *Link) AddToTail(value int) {
	currNode := l.head
	fmt.Println("head: ", l.head)
	newData := Node{}
	newData.data = value
	if l.head == nil {
		l.head = &newData
		l.len++
	} else {
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = &newData
		l.len++
	}
}

// add to head/ prepend
func (l *Link) Prepend(value int) {
	newNode := Node{}
	newNode.data = value
	currNode := l.head
	// special case
	if l.head == nil {
		l.head = &newNode
	}
	if l.head != nil {
		l.head = &newNode
		l.head.next = currNode
		l.len++
	}

}

// traverse all the value
func (l *Link) ShowAll() {
	currNode := l.head
	fmt.Printf("%v -> ", currNode.data)
	for currNode.next != nil {
		currNode = currNode.next
		fmt.Printf("%v -> ", currNode.data)
	}
}

// next need to learn using INTERFACE

func main() {
	link := Link{}
	link.AddToTail(12)
	link.AddToTail(13)
	link.Prepend(50)
	link.AddToTail(15)
	link.Prepend(500)
	link.ShowAll()
}
