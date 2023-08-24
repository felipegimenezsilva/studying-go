package main

import "fmt"

// insert a new node in the tree, or return the node if empty tree
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

// Print all the values in the tree in Ascending order
func PrintOrderRecursive(root *NodeBinaryTree) {
	if root != nil {
		PrintOrderRecursive((*root).right)
		fmt.Printf("%v ", (*root).value)
		PrintOrderRecursive((*root).left)
	}
}
