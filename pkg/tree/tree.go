package tree

import "fmt"

type TreeNode struct {
	Name    string
	MinTime uint
	MaxTime uint
	gmax    uint
	height  uint
	left    *TreeNode
	right   *TreeNode
}

type Tree struct {
	root *TreeNode
}

func NewTree() *Tree {
	return &Tree{}
}

func (tn *TreeNode) isOverlapping(qmin uint, qmax uint) bool {

	if tn.MaxTime <= qmin || tn.MinTime >= qmax {
		return false
	}

	return true
}

func (t *Tree) Insert(name string, minTime uint, maxTime uint) {

	n := &TreeNode{Name: name, MinTime: minTime, MaxTime: maxTime, gmax: maxTime}
	if t.root == nil {
		t.root = n
		return
	}
	// insertNode(t.root, n)

	t.root = balancedInsert(t.root, n)
}

func searchOverlaps(node *TreeNode, qmin uint, qmax uint) []*TreeNode {
	overlaps := make([]*TreeNode, 0)
	if node.isOverlapping(qmin, qmax) {
		overlaps = append(overlaps, node)
	}

	if node.left != nil && node.left.gmax > qmin {
		overlaps = append(overlaps, searchOverlaps(node.left, qmin, qmax)...)
	}

	if node.right != nil && node.right.gmax > qmin {
		overlaps = append(overlaps, searchOverlaps(node.right, qmin, qmax)...)
	}
	return overlaps
}

func (t *Tree) Search(qmin uint, qmax uint) []*TreeNode {
	if t.root == nil {
		return []*TreeNode{}
	}

	return searchOverlaps(t.root, qmin, qmax)
}

func (t *TreeNode) InorderTraversal() {
	if t == nil {
		return
	}
	t.left.InorderTraversal()
	fmt.Printf("\n%+v", t)
	t.right.InorderTraversal()
}
