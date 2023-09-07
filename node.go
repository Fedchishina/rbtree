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

func isRed[V constraints.Ordered](n *node[V]) bool {
	return n.color == red
}
func isBlack[V constraints.Ordered](n *node[V]) bool {
	return n.color == black
}

func isLeftChild[V constraints.Ordered](n *node[V]) bool {
	return n == n.parent.left
}

func isRightChild[V constraints.Ordered](n *node[V]) bool {
	return n == n.parent.right
}

func recolorForInsertCase1[V constraints.Ordered](y, z *node[V]) {
	z.parent.color = black
	y.color = black
	z.parent.parent.color = red
}

func recolorForInsertCase3[V constraints.Ordered](z *node[V]) {
	z.parent.color = black
	z.parent.parent.color = red
}
