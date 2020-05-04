package binarytree

import "os"

// BinaryTree ...
type BinaryTree struct {
	Root *BinaryNode
}

// NewBinaryTree ...
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

// Insert add node to the tree
func (t *BinaryTree) Insert(key float64, data interface{}) *BinaryTree {
	if t.Root == nil {
		t.Root = NewBinaryNode(key, data)
	} else {
		t.Root.Insert(key, data)
	}
	return t
}

// Max returns the max value node
func (t *BinaryTree) Max() *BinaryNode {
	n := t.Root
	if n == nil {
		return nil
	}
	for {
		if n.Right == nil {
			return n
		}
		n = n.Right
	}
}

// Print displays tree from root
func (t *BinaryTree) Print() {
	t.Root.Print(os.Stdout, 0, 'M')
}
