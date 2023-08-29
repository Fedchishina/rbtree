package rbtree

import "golang.org/x/exp/constraints"

type color bool

const (
	red   color = true
	black color = false
)

// node is the structure of tree's node.
// node's key is any ordered type for type of
// node's value has type any
type node[V constraints.Ordered] struct {
	element element[V]
	parent  *node[V]
	left    *node[V]
	right   *node[V]
	color   color
}

type element[V constraints.Ordered] struct {
	key   V
	value any
}

func search[V constraints.Ordered](n *node[V], key V) *node[V] {
	for n != nil && key != n.element.key {
		if key < n.element.key {
			n = n.left
			continue
		}
		n = n.right
	}

	return n
}

func (n *node[V]) hasNoChildren() bool {
	return n.left == nil && n.right == nil
}

func rightRotate[V constraints.Ordered](root, y *node[V]) {
	x := y.left
	y.left = x.right

	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if y.parent == nil {
		root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}

	x.right = y
	y.parent = x
}
