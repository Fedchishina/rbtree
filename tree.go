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

// Delete is a function for deleting node in rbtree
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) Delete(key V) {
	z := search(t.root, key)
	if z == nil {
		return
	}

	if t.root.hasNoChildren() {
		t.root = nil
		return
	}

	yOriginalColor, x := t.deleteNode(z)

	if yOriginalColor == black {
		t.deleteFixup(x)
	}
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
		if v != nil {
			v.parent = nil
		}

		return
	}

	if u == u.parent.left {
		u.parent.left = v
		v.parent = u.parent

		return
	}

	u.parent.right = v

	if v != nil {
		v.parent = u.parent
	}
}

// deleteNode - internal function for deleting node in rbtree
func (t *Tree[V]) deleteNode(z *node[V]) (color, *node[V]) {
	var yOriginalColor color
	y := z
	yOriginalColor = y.color

	var x *node[V]

	if z.left == nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = z.right.min()
		yOriginalColor = y.color
		x = y.right

		if y.parent != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	return yOriginalColor, x
}

func (t *Tree[V]) deleteFixup(x *node[V]) {
	var w *node[V]
	for x != t.root && x.color == black {
		if x == x.parent.left {
			w = x.parent.right
			if w.color == red {
				w.color = black
				x.parent.color = red
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else if w.right.color == black {
				w.left.color = black
				w.color = red
				t.rightRotate(w)
				w = x.parent.right
			}
			w.color = x.parent.color
			x.parent.color = black
			w.right.color = black
			t.leftRotate(x.parent)
			x = t.root
		} else {
			w = x.parent.left
			if w.color == red {
				w.color = black
				x.parent.color = red
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.parent
			} else if w.left.color == black {
				w.right.color = black
				w.color = red
				t.leftRotate(w)
				w = x.parent.left
			}
			w.color = x.parent.color
			x.parent.color = black
			w.left.color = black
			t.rightRotate(x.parent)
			x = t.root
		}
	}
	x.color = black
}
