package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

// type Node3 struct {
// 	Value  int
// 	Left   *Node3
// 	Center *Node3
// 	Right  *Node3
// }

// func InsertNode(node *Node3, val int) *Node3 {
// 	NewNode := &Node3{Value: val}
// 	if node == nil {
// 		return NewNode
// 	}
// 	if node.Left == nil {
// 		node.Left = InsertNode(node.Left, val)
// 	} else if node.Center == nil {
// 		node.Center = InsertNode(node.Center, val)
// 	} else if node.Right == nil {
// 		node.Right = InsertNode(node.Right, val)
// 	}
// 	return node
// }
// func InoderTraversal3(root *Node3) {
// 	if root != nil {
// 		InoderTraversal3(root.Left)
// 		InoderTraversal3(root.Center)
// 		fmt.Println(root.Value)
// 		InoderTraversal3(root.Right)
// 	}
// }
// func main() {
// 	root := &Node3{}
// 	new := []int{10, 11, 5, 2, 6, 8}
// 	for _, v := range new {
// 		root = InsertNode(root, v)
// 	}
// 	InoderTraversal3(root)

// }

// func (bst *BinarySearchTree) Insert(value int) {
// 	newNode := &Node{Value: value}
// 	if bst.Root == nil {
// 		bst.Root = newNode
// 	} else {
// 		bst.Root.insertNode(newNode)
// 	}
// }

//	func (n *Node) insertNode(newNode *Node) {
//		if newNode.Value < n.Value {
//			if n.Left == nil {
//				n.Left = newNode
//			} else {
//				n.Left.insertNode(newNode)
//			}
//		} else {
//			if n.Right == nil {
//				n.Right = newNode
//			} else {
//				n.Right.insertNode(newNode)
//			}
//		}
//	}
func Search(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if root.Value > value {
		return Search(root.Left, value)
	} else if root.Value < value {
		return Search(root.Right, value)
	}
	return true
}
func DeleteNode(n *Node, value int) *Node {
	if n == nil {
		return nil
	}
	if n.Value > value {
		n.Left = DeleteNode(n.Left, value)
	} else if n.Value < value {
		n.Right = DeleteNode(n.Right, value)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}
		MinRight := FindMin(n.Right)
		n.Value = MinRight.Value
		n.Right = DeleteNode(n.Right, MinRight.Value)
	}
	return n
}
func FindMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
func Height(node *Node) int {
	if node == nil {
		return 0
	}
	left := Height(node.Left)
	right := Height(node.Right)
	return max(left, right) + 1
}
func balanced(node *Node) bool {
	if node == nil {
		return true
	}
	left := Height(node.Left)
	right := Height(node.Right)
	if abs(left-right) > 1 {
		return false
	}
	return true
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func BFS(root *Node) {
	if root == nil {
		return
	}

	// Create a queue to hold nodes
	queue := []*Node{}
	// Start with the root node
	queue = append(queue, root)

	for len(queue) > 0 {
		// Dequeue the front node
		current := queue[0]
		queue = queue[1:]

		// Process the current node (e.g., print its value)
		fmt.Printf("%d ", current.Value)

		// Enqueue the left child if it exists
		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		// Enqueue the right child if it exists
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func main() {
	var root *Node
	new := []int{10, 11, 5, 2, 6, 8}
	for _, v := range new {
		root = insertion(root, v)
	}
	DFS(root)
	// fmt.Println()
	InoderTraversal(root)
	// fmt.Println("height", Height(bst.Root))
	// DeleteNode(bst.Root, 3)
	// Preorder(root)
	// if balanced(bst.Root) {
	// 	fmt.Println("balanced")
	// } else {
	// 	fmt.Println("unbalanced")
	// }
	// fmt.Println("Search 7:", Search(root, 33))
	// fmt.Println("Search 20:", bst.Search(20))
}
func InoderTraversal(root *Node) {
	if root != nil {
		InoderTraversal(root.Left)
		fmt.Println(root.Value)
		InoderTraversal(root.Right)
	}
}
func Preorder(node *Node) {
	if node != nil {
		fmt.Println(node.Value)
		Preorder(node.Left)
		Preorder(node.Right)
	}
}
func Postorder(node *Node) {
	if node != nil {
		Preorder(node.Left)
		Preorder(node.Right)
		fmt.Println(node.Value)
	}
}

func DFS(root *Node) {
	fmt.Println(root.Value)
	if root.Left != nil {
		DFS(root.Left)
	} else if root.Right != nil {
		DFS(root.Right)
	}
}
func insertion(root *Node, value int) *Node {
	newNode := &Node{Value: value}
	if root == nil {
		return newNode
	}
	if root.Value > value {
		root.Left = insertion(root.Left, value)
	} else if root.Value < value {
		root.Right = insertion(root.Right, value)
	}
	return root
}
