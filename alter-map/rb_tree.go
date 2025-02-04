package rbtree

// RbTree is a self balancing BST that is used as an alternative to hashmap ds.
// It supports all operations that a map does but more efficiently.
type RbTree struct {
	Length int

	root *node
	// iter *iterator TODO: implement in order iterator
}

func New() *RbTree {
	return &RbTree{}
}

func (t *RbTree) Get(k int) any {
	current := t.root

	for current != nil {
		if k == current.key {
			return current.value
		} else if k < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	return nil
}

// Put updates the value of key k to value v. New entry is created if key is not present.
func (t *RbTree) Put(k int, v any) {
	if t.root == nil {
		t.root = &node{
			key:       k,
			value:     v,
			nodeColor: black,
		}
		t.Length = 1
		return
	}

	var (
		current *node = t.root
		parent  *node = nil
		exists  bool  = false
	)

	for current != nil && !exists {
		if k == current.key {
			current.value = v
			exists = true
		} else if k > current.key {
			parent = current
			current = current.right
		} else {
			parent = current
			current = current.left
		}
	}

	if !exists {
		newNode := &node{
			key:    k,
			value:  v,
			parent: parent,
		}

		if k > parent.key {
			parent.right = newNode
		} else {
			parent.left = newNode
		}

		t.Length++
		t.fixInsert(newNode)
	}
}

func (t *RbTree) fixInsert(n *node) {
	for n.parent != nil && n.parent.nodeColor == red {
		if n.parent == n.parent.parent.right {
			uncle := n.parent.parent.left
			if uncle != nil && uncle.nodeColor == red {
				// Case 1: Uncle is red
				uncle.nodeColor = black
				n.parent.nodeColor = black
				n.parent.parent.nodeColor = red
				n = n.parent.parent
			} else {
				// Case 2: Uncle is black or nil
				if n == n.parent.left {
					// Case 2a: Left child of right child
					n = n.parent
					n.rotateRight(t)
				}
				// Case 2b: Right child of right child
				n.parent.nodeColor = black
				n.parent.parent.nodeColor = red
				n.parent.parent.rotateLeft(t)
			}
		} else {
			uncle := n.parent.parent.right
			if uncle != nil && uncle.nodeColor == red {
				// Case 1: Uncle is red (symmetric)
				uncle.nodeColor = black
				n.parent.nodeColor = black
				n.parent.parent.nodeColor = red
				n = n.parent.parent
			} else {
				// Case 2: Uncle is black or nil (symmetric)
				if n == n.parent.right {
					// Case 2a: Right child of left child
					n = n.parent
					n.rotateLeft(t)
				}
				// Case 2b: Left child of left child
				n.parent.nodeColor = black
				n.parent.parent.nodeColor = red
				n.parent.parent.rotateRight(t)
			}
		}
		if n == t.root {
			break
		}
	}

	// recolor root to black if it was changed during fixup
	t.root.nodeColor = black
}
