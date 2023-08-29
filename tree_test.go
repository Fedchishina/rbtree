package rbtree

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestNew(t *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		want *Tree[V]
	}
	testInt := testCase[int]{
		name: "int empty tree",
		want: &Tree[int]{root: nil},
	}
	t.Run(testInt.name, func(t *testing.T) {
		if got := New[int](); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})

	testString := testCase[string]{
		name: "int empty tree",
		want: &Tree[string]{root: nil},
	}
	t.Run(testString.name, func(t *testing.T) {
		if got := New[int](); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})
}

func TestNewWithElement(t *testing.T) {
	type args[V constraints.Ordered] struct {
		key   V
		value any
	}
	type testCase[V constraints.Ordered] struct {
		name string
		args args[V]
		want *Tree[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty value",
			args: args[int]{key: 1, value: nil},
			want: &Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   1,
						value: nil,
					},
					color: black,
				},
			},
		},
		{
			name: "one element",
			args: args[int]{key: 15, value: 15},
			want: getTree([]int{15}),
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithElement(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Min(t1 *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		want V
	}

	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			want: 0,
		},
		{
			name: "tree with one element",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			want: 15,
		},
		{
			name: "tree with root and one element",
			t:    *treeWithOneElement,
			want: 15,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Min(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Max(t1 *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		want V
	}

	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			want: 0,
		},
		{
			name: "tree with one element",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			want: 15,
		},
		{
			name: "tree with root and one element",
			t:    *treeWithOneElement,
			want: 25,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Max(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Exist(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want bool
	}
	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			args: args[int]{key: 1},
			want: false,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 1},
			want: false,
		},
		{
			name: "tree with one element - found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 15},
			want: true,
		},
		{
			name: "tree with root and one element - found",
			t:    *treeWithOneElement,
			args: args[int]{key: 25},
			want: true,
		},
		{
			name: "tree with root and one element - not found",
			t:    *treeWithOneElement,
			args: args[int]{key: 35},
			want: false,
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Exists(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_GetValue(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name    string
		t       Tree[V]
		args    args[V]
		want    any
		wantErr bool
	}
	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name:    "empty tree",
			t:       Tree[int]{root: nil},
			args:    args[int]{key: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args:    args[int]{key: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tree with one element - found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args:    args[int]{key: 15},
			want:    15,
			wantErr: false,
		},
		{
			name:    "tree with root and one element - found",
			t:       *treeWithOneElement,
			args:    args[int]{key: 25},
			want:    25,
			wantErr: false,
		},
		{
			name:    "tree with root and one element - not found",
			t:       *treeWithOneElement,
			args:    args[int]{key: 35},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.t.GetValue(tt.args.key)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Insert(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key   V
		value any
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    *Tree[V]
		args args[V]
		want *Tree[V]
	}

	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
				color:  red,
			},
			color: black,
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    &Tree[int]{},
			args: args[int]{key: 15, value: 15},
			want: &Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
					color: black,
				},
			},
		},
		{
			name: "tree with root and one element",
			t:    getTree([]int{15}),
			args: args[int]{key: 25, value: 25},
			want: &treeWithOneElement,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Insert(tt.args.key, tt.args.value)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Insert() = %#+v, want %#+v", tt.t, tt.want)
			}
		})
	}
}

func TestTree_Insert_big_case(t1 *testing.T) {
	t := getTree([]int{11, 2, 14, 1, 7, 15, 5, 8})

	// check tree's structure and colours before insert
	checkNodeProperties(t1, t.root, 11, black, "t.root")
	checkNodeProperties(t1, t.root.left, 2, red, "t.root.left")
	checkNodeProperties(t1, t.root.right, 14, black, "t.root.right")
	checkNodeProperties(t1, t.root.right.right, 15, red, "t.root.right.right")
	checkNodeProperties(t1, t.root.left.left, 1, black, "t.root.left.left")
	checkNodeProperties(t1, t.root.left.right, 7, black, "t.root.left.right")
	checkNodeProperties(t1, t.root.left.right.left, 5, red, "t.root.left.right.left")
	checkNodeProperties(t1, t.root.left.right.right, 8, red, "t.root.left.right.right")

	// add list
	t.Insert(4, 4)

	// check tree's structure and colours after insert
	checkNodeProperties(t1, t.root, 7, black, "t.root")
	checkNodeProperties(t1, t.root.left, 2, red, "t.root.left")
	checkNodeProperties(t1, t.root.right, 11, red, "t.root.right")
	checkNodeProperties(t1, t.root.left.left, 1, black, "t.root.left.left")
	checkNodeProperties(t1, t.root.left.right, 5, black, "t.root.left.right")
	checkNodeProperties(t1, t.root.right.left, 8, black, "t.root.right.left")
	checkNodeProperties(t1, t.root.right.right, 14, black, "t.root.right.right")
	checkNodeProperties(t1, t.root.right.right.right, 15, red, "t.root.right.right.right ")
	checkNodeProperties(t1, t.root.left.right.left, 4, red, "t.root.left.right.left")
}

func checkNodeProperties(t *testing.T, node *node[int], key int, color color, errMsg string) {
	if node == nil {
		return
	}

	if node.element.key != key || node.color != color {
		t.Errorf("Error - Want key: %v, have key %v. Want color: %v, have color: %v in %s",
			key,
			node.element.key,
			color,
			node.color,
			errMsg,
		)
	}
}

func getTree(elements []int) *Tree[int] {
	tree := New[int]()
	for _, el := range elements {
		tree.Insert(el, el)
	}

	return tree
}
