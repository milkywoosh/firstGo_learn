package main
import ("fmt")

type Node struct {
	element int
	next *Node
}
type Link struct {
	head *Node
	len int
}
func (l *Link)AddToTail(value int){
	currNode := l.head
	//
	newNode := &Node{}
	newNode.element = value
	if l.head == nil {
		l.head = newNode
		l.len ++
	} else {
		// form of WHILE in GOLANG
		
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = newNode
		l.len ++
	}
}
// add data from index 0

func (l *Link) Prepend(value int){
	currNode := l.head
	newNode := Node{}
	newNode.element = value
	if l.head == nil && l.len == 0 {
		l.head = &newNode
		l.len ++
	} else {
		helper := &newNode // {element: value, next: nil}
		helper.next = currNode
		//{element: next: l.head}
		l.head = helper
		l.len++
	} 
	
}
// show backwards


func (l *Link)ShowAll(){
	currNode := l.head

	// fmt.Printf("%v -> ", currNode.element)
	for currNode != nil {
		fmt.Printf("%v -> ", currNode.element)
		currNode = currNode.next
	}
	fmt.Println()
}
func (l *Link)Length() {
	fmt.Println("\nlen of data: ", l.len)
}

func main() {
	link := Link{}
	data := []int{2,3,6,5,3,8,1,12}
	for _, value := range data {
		link.AddToTail(value)
	}
	link.Prepend(30)
	link.Prepend(37)
	link.ShowAll()
	link.Length()

}