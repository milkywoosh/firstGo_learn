// binary tree
/*
		    root
		    /  \
		c_node	 c_node
		/  \           \
	  c_node   c_node	 node
*/
package main

import "fmt"

// Node represent components
type Node struct {
	Root  int
	Left  *Node // is type of Node
	Right *Node // is type of Node -> *Node
}

// Insert(data)
// asterisk is for directly edit data in struct??
func (n *Node) Insert(val int) {
	if val > n.Root {
		// checking to right branch
		if n.Right == nil {
			n.Right = &Node{Root: val}
		} else if n.Right != nil {
			n.Right.Insert(val) // this is recursion ??
		}
	} else if val < n.Root {
		// checking to left branch
		if n.Left == nil {
			n.Left = &Node{Root: val}
		} else if n.Left != nil {
			n.Left.Insert(val) // this is recursion ??
		}
	}
}

// search(data)
func (n *Node) Search(val int) bool {

	if n == nil {
		return false
	}
	// why gabisa if val == n.Root here???

	// first check  is val is bigger than ROOT
	if val > n.Root {
		// cek reCURSION
		fmt.Println(n.Root)
		return n.Right.Search(val) // FORM OF RECURSION TO CHECK RIGHT CHILD
		// else check this line
	} else if val < n.Root {
		return n.Left.Search(val)
	}
	fmt.Println(val)

	// go ALWAYS NEED LAST RETURN behaviour???
	return val == n.Root

}

// leastVal()
// greatestVal()

func main() {
	// struct literal
	// tree here as an "address"
	tree := &Node{Root: 100}
	tree.Insert(90)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(400)
	fmt.Println(tree)

	fmt.Println(tree.Search(400))
}
