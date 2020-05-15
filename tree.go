package binarytree

import "os"

// BinaryTree ...
type BinaryTree struct {
	Root          *BinaryNode
	activateSplay bool
}

// NewBinaryTree ...
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil, activateSplay: false}
}

// ToggleSplay toggle to use splay or not
func (t *BinaryTree) ToggleSplay(activate bool) {
	t.activateSplay = activate
}

// Insert add node to the tree
func (t *BinaryTree) Insert(key float64, data interface{}) *BinaryTree {
	if t.Root == nil {
		t.Root = NewBinaryNode(key, data)
	} else {
		t.Root.Insert(key, data)
		if t.activateSplay {
			t.splay(key)
		}
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
			if t.activateSplay {
				t.splay(n.Key)
			}
			return n
		}
		n = n.Right
	}
}

// Min returns the min value node
func (t *BinaryTree) Min() *BinaryNode {
	n := t.Root
	if n == nil {
		return nil
	}
	for {
		if n.Left == nil {
			if t.activateSplay {
				t.splay(n.Key)
			}
			return n
		}
		n = n.Left
	}
}

// Print displays tree from root
func (t *BinaryTree) Print() {
	t.Root.Print(os.Stdout, 0, 'M')
}
