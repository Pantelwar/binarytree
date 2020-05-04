package binarytree

import (
	"fmt"
	"io"
)

// type Item generic.Type

// BinaryNode ...
type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Key   float64
	Data  interface{}
}

// NewBinaryNode returns a new node
func NewBinaryNode(key float64, data interface{}) *BinaryNode {
	return &BinaryNode{Key: key, Data: data, Left: nil, Right: nil}
}

// Insert adds a new node to the tree
func (n *BinaryNode) Insert(key float64, data interface{}) {
	if n == nil {
		return
	} else if key < n.Key {
		// fmt.Println(" key < n.Key", key, n.Key)
		if n.Left == nil {
			n.Left = NewBinaryNode(key, data)
		} else {
			n.Left.Insert(key, data)
		}
	} else if key > n.Key {
		// fmt.Println(" key > n.Key", key, n.Key)
		if n.Right == nil {
			n.Right = NewBinaryNode(key, data)
		} else {
			n.Right.Insert(key, data)
		}
	} else {
		// fmt.Println(" key = n.Key", key, n.Key)

	}
}

// Print displays tree from the node
func (n *BinaryNode) Print(w io.Writer, ns int, ch rune) {
	if n == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v -> %#v\n", ch, n.Key, n.Data)
	n.Left.Print(w, ns+2, 'L')
	n.Right.Print(w, ns+2, 'R')
}
