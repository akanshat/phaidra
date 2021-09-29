package tree

// InsertNode is naive binary search insertion.
// This is not used now.
// Instead we use self-balancing insertion similar to AVL-trees.
func InsertNode(node *TreeNode, newNode *TreeNode) {

	if node.MinTime > newNode.MinTime {
		if node.left == nil {
			node.left = newNode
		} else {
			InsertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			InsertNode(node.right, newNode)
		}
	}
	if newNode.MaxTime > node.gmax {
		node.gmax = newNode.MaxTime
	}
}
