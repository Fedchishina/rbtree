package rbtree

import "golang.org/x/exp/constraints"

type color int

const (
	red color = iota
	black
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
