package main
import ("fmt")

type Node struct {
	value int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}
// if in this func it use * then it should modify the struct
func (linkedList *LinkedList) AddToHead(value int) {
	var node = Node{}
	node.value = value
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.value)
	}

}

func main() {
	var li LinkedList
	li = LinkedList{}
	li.AddToHead(3)
	li.AddToHead(4)
	li.AddToHead(5)

	// iterateList
	li.IterateList()
	
}