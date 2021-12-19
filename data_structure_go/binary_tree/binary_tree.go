// binary tree
/*
			root
		node	node
	node			node
*/
package main
import "fmt"

// Node represent components
type Node struct {
	Root int
	Left *Node
	Right *Node
}

// Insert(data)
func (n *Node) Insert(val int) {
	if val > n.Root {
		// checking to right branch
		if n.Right == nil {
			n.Right = &Node{Root: val}
		} else if n.Right != nil {
			n.Right.Insert(val)
		}
	} else if val < n.Root {
		// checking to left branch
		if n.Left == nil {
			n.Left = &Node{Root: val}
		} else if n.Left != nil {
			n.Left.Insert(val)
		}
	}
}
// search(data)
func (n *Node) Search(val int) bool {

	if n == nil {
		return false
	}
	// why gabisa if val == n.Root here???
	if val > n.Root {
		return n.Right.Search(val)
	} else if val < n.Root {
		return n.Left.Search(val)
	}
	fmt.Println(val)
	return val == n.Root 
	
}
// leastVal()
// greatestVal()

func main() {
	// struct literal 
	// tree here as an "address"
	tree := &Node{Root:100}
	tree.Insert(90)
	tree.Insert(200)
	fmt.Println(tree)

	fmt.Println(tree.Search(90))
}








