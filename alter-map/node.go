package rbtree

type color byte

const (
	red   color = 0
	black color = 1
)

type node struct {
	key       int
	value     any
	right     *node
	left      *node
	parent    *node
	nodeColor color
}

func (n *node) rotateLeft(tree *RbTree) {
	right := n.right
	n.right = right.left

	if right.left != nil {
		right.left.parent = n
	}

	right.parent = n.parent

	if n.parent == nil {
		tree.root = right
	} else if n == n.parent.left {
		n.parent.left = right
	} else {
		n.parent.right = right
	}

	right.left = n
	n.parent = right
}

func (n *node) rotateRight(tree *RbTree) {
	left := n.left
	n.left = left.right

	if left.right != nil {
		left.right.parent = n
	}

	left.parent = n.parent

	if n.parent == nil {
		tree.root = left
	} else if n == n.parent.left {
		n.parent.left = left
	} else {
		n.parent.right = left
	}

	left.right = n
	n.parent = left
}
