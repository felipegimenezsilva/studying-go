package main

import "fmt"

type NodeBinaryTree struct {
	value int
	left  *NodeBinaryTree
	right *NodeBinaryTree
}

func InsertNodeRecursive(root *NodeBinaryTree, node *NodeBinaryTree) *NodeBinaryTree {
	if root == nil {
		return node
	} else if root.value > node.value {
		if (*root).right != nil {
			InsertNodeRecursive((*root).right, node)
		} else {
			(*root).right = node
		}
	} else {
		if (*root).left != nil {
			InsertNodeRecursive((*root).left, node)
		} else {
			(*root).left = node
		}
	}
	return root
}

// Return Pointers of Local Variables Is Safe in Go:
// Unlike C language, Go is a language supporting garbage collection,
// so return the address of a local variable is absolutely safe in Go.
// --> https://go101.org/article/pointer.html
// ## Go makes some restrictions to pointers (comparing to pointers in C language).
func CreateLeaf(value int) *NodeBinaryTree {
	var node NodeBinaryTree
	node.value = value
	node.left = nil
	node.right = nil
	return &node
}

func main() {

	var root *NodeBinaryTree = nil

	root = InsertNodeRecursive(root, CreateLeaf(7))
	root = InsertNodeRecursive(root, CreateLeaf(8))

	fmt.Printf("%v", *root)

}
