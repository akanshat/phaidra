package tree

import (
	"math"
)

func height(t *TreeNode) uint {
	if t == nil {
		return 0
	}
	return t.height
}

func max(left uint, right uint) uint {
	return uint(math.Max(float64(left), float64(right)))
}

func getBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return int(height(node.left)) - int(height(node.right))
}

func rightRotate(node *TreeNode) *TreeNode {
	x := node.left
	r := x.right

	node.gmax = max(node.gmax, r.gmax)
	x.right = node
	x.gmax = max(x.gmax, node.gmax)
	node.left = r

	node.height = 1 + max(height(node.left), height(node.right))
	x.height = 1 + max(height(x.left), height(x.right))

	return x
}

func leftRotate(node *TreeNode) *TreeNode {
	y := node.right
	l := y.left

	node.gmax = max(node.gmax, l.gmax)
	y.left = node
	y.gmax = max(y.gmax, node.gmax)
	node.right = l

	node.height = 1 + max(height(node.left), height(node.right))
	y.height = 1 + max(height(y.left), height(y.right))

	return y
}

func balancedInsert(node *TreeNode, newNode *TreeNode) *TreeNode {

	if node.MinTime > newNode.MinTime {
		if node.left == nil {
			node.left = newNode
			return node
		} else {
			node.left = balancedInsert(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
			return node
		} else {
			node.right = balancedInsert(node.right, newNode)
		}
	}
	if newNode.MaxTime > node.gmax {
		node.gmax = newNode.MaxTime
	}

	node.height = 1 + max(height(node.left), height(node.right))

	balance := getBalance(node)

	// Left-left case.
	if balance > 1 && newNode.MinTime < node.left.MinTime {
		return rightRotate(node)
	}

	// Right-right case.
	if balance < -1 && newNode.MinTime > node.right.MinTime {
		return leftRotate(node)
	}

	// Left-right case.
	if balance > 1 && newNode.MinTime > node.left.MinTime {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// Right-left case.
	if balance < -1 && newNode.MinTime < node.right.MinTime {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node

}
