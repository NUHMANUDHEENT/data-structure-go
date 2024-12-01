package main

import "fmt"

// TreeNode represents a node in the tree with n children.
type TreeNode struct {
	value    int
	children []*TreeNode
}

// NewTreeNode creates a new tree node.
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		value:    value,
		children: []*TreeNode{},
	}
}

// AddChild adds a new child to the current node.
func (n *TreeNode) AddChild(child *TreeNode) {
	n.children = append(n.children, child)
}

// PrintTree recursively prints the tree in a structured format.
func (n *TreeNode) PrintTree() {
	// Print the value of the current node with indentation for levels
	fmt.Print(n.value)
	// Recursively print the children
	for _, child := range n.children {
		child.PrintTree()
	}
}


func main() {
	// Create the root of the tree
	root := NewTreeNode(1)

	// Add children to the root node
	child1 := NewTreeNode(2)
	child2 := NewTreeNode(3)
	root.AddChild(child1)
	root.AddChild(child2)

	// Add children to the first child of the root
	grandChild1 := NewTreeNode(4)
	grandChild2 := NewTreeNode(5)
	child1.AddChild(grandChild1)
	child1.AddChild(grandChild2)

	// Add children to the second child of the root
	grandChild3 := NewTreeNode(6)
	child2.AddChild(grandChild3)

	// Print the entire tree
	fmt.Println("Tree structure:")
	root.PrintTree()
}
