package binarytree

// Splay functionality
func (tree *BinaryTree) splay(key float64) {
	var temp = BinaryNode{}
	var lef = &temp
	var rig = &temp
	var top = tree.Root
	for {
		if key < top.Key { //item.Less(top.item) {
			if top.Left == nil {
				break
			}
			if key < top.Left.Key { //item.Less(top.left.item) {
				// rotate right
				yes := top.Left
				top.Left = yes.Right
				yes.Right = top
				top = yes
				if top.Left == nil {
					break
				}
			}
			// link right
			rig.Left = top
			rig = top
			top = top.Left
		} else if top.Key < key { //top.item.Less(item) {
			if top.Right == nil {
				break
			}
			if top.Right.Key < key { // top.right.item.Less(item) {
				// rotate left
				yes := top.Right
				top.Right = yes.Left
				yes.Left = top
				top = yes
				if top.Right == nil {
					break
				}
			}
			// link left
			lef.Right = top
			lef = top
			top = top.Right
		} else {
			break
		}
	}
	// assemble
	lef.Right = top.Left
	rig.Left = top.Right
	top.Left = temp.Right
	top.Right = temp.Left
	tree.Root = top
}
