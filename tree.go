package rbtree

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Tree[V constraints.Ordered] struct {
	root *node[V]
}

// New is a function for creation empty tree
// - param should be `ordered type` (`int`, `string`, `float` etc)
func New[V constraints.Ordered]() *Tree[V] {
	return &Tree[V]{}
}

// NewWithElement is a function for creation tree with one element
// - param should be `ordered type` (`int`, `string`, `float` etc)
func NewWithElement[V constraints.Ordered](key V, value any) *Tree[V] {
	return &Tree[V]{
		root: &node[V]{
			element: element[V]{
				key:   key,
				value: value,
			},
			color: black,
		},
	}
}

// Insert is a function for inserting element into Tree
// - param key should be `ordered type` (`int`, `string`, `float` etc.)
// - param value can be any type
func (t *Tree[V]) Insert(key V, value any) {
	newNode := &node[V]{element: element[V]{
		key:   key,
		value: value,
	},
		color: red,
	}

	if t.root == nil {
		t.root = newNode
		t.root.color = black

		return
	}

	current := t.root
	for {
		if key < current.element.key {
			if current.left == nil {
				current.left = newNode
				newNode.parent = current
				break
			}
			current = current.left
			continue
		}

		if current.right == nil {
			current.right = newNode
			newNode.parent = current
			break
		}
		current = current.right
	}
	t.insertFixup(newNode)
}

// Min is a function for searching min element in tree (by key).
func (t *Tree[V]) Min() V {
	var result V
	n := t.root
	if n == nil {
		return result
	}

	for n.left != nil {
		n = n.left
	}

	return n.element.key
}

// Max is a function for searching max element in tree (by key).
func (t *Tree[V]) Max() V {
	var result V
	n := t.root
	if n == nil {
		return result
	}

	for n.right != nil {
		n = n.right
	}

	return n.element.key
}

// Exists is a function for searching element in node. If element exists in tree - return true, else - false
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) Exists(key V) bool {
	searchNode := search(t.root, key)
	if searchNode == nil {
		return false
	}

	return true
}

// GetValue is a function for searching element in node and returning value of this element
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) GetValue(key V) (any, error) {
	var result any
	searchNode := search(t.root, key)
	if searchNode == nil {
		return result, errors.New(fmt.Sprintf("element with key %v not found", key))
	}

	return searchNode.element.value, nil
}

// leftRotate - internal function for left rotating in rbtree
func (t *Tree[V]) leftRotate(x *node[V]) {
	if x == nil || x.right == nil {
		return
	}

	y := x.right
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		t.root = y
		y.left = x
		x.parent = y

		return
	}

	if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

// rightRotate - internal function for right rotating in rbtree
func (t *Tree[V]) rightRotate(y *node[V]) {
	if y == nil || y.left == nil {
		return
	}

	x := y.left
	y.left = x.right

	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}

	x.right = y
	y.parent = x
}

// insertFixup function calls after insert node to rbtree for recovery of rbtree's properties
func (t *Tree[V]) insertFixup(z *node[V]) {
	for z.parent != nil && z.parent.color == red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y != nil && y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
				continue
			}
			if z == z.parent.right {
				z = z.parent
				t.leftRotate(z)
			}
			z.parent.color = black
			z.parent.parent.color = red
			t.rightRotate(z.parent.parent)
			continue
		}
		y := z.parent.parent.left
		if y != nil && y.color == red {
			z.parent.color = black
			y.color = black
			z.parent.parent.color = red
			z = z.parent.parent
			continue
		}

		if z == z.parent.left {
			z = z.parent
			t.rightRotate(z)
		}
		z.parent.color = black
		z.parent.parent.color = red
		t.leftRotate(z.parent.parent)
	}
	t.root.color = black
}

// transplant - internal function for substitution u node to v node
func (t *Tree[V]) transplant(u, v *node[V]) {
	if u.parent == nil {
		t.root = v

		return
	}

	if u == u.parent.left {
		u.parent.left = v
		v.parent = u.parent

		return
	}

	u.parent.right = v
	v.parent = u.parent
}
