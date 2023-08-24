package main

import "fmt"

type NodeBinaryTree struct {
	value int
	left  *NodeBinaryTree
	right *NodeBinaryTree
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

	root = InsertNodeRecursive(root, CreateLeaf(5))
	root = InsertNodeRecursive(root, CreateLeaf(8))
	root = InsertNodeRecursive(root, CreateLeaf(1))
	root = InsertNodeRecursive(root, CreateLeaf(8))
	root = InsertNodeRecursive(root, CreateLeaf(4))
	root = InsertNodeRecursive(root, CreateLeaf(7))
	root = InsertNodeRecursive(root, CreateLeaf(6))
	root = InsertNodeRecursive(root, CreateLeaf(9))
	root = InsertNodeRecursive(root, CreateLeaf(10))
	root = InsertNodeRecursive(root, CreateLeaf(2))
	root = InsertNodeRecursive(root, CreateLeaf(3))
	root = InsertNodeRecursive(root, CreateLeaf(0))

	PrintOrderRecursive(root)
	fmt.Printf("\n")
	PrintTreeOrderRecursive(root)

}
