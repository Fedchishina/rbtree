package rbtree

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

type validNode[V constraints.Ordered] struct {
	node     *node[V]
	key      V
	color    color
	nodePath string
}

func TestTree_New(t *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		want *Tree[V]
	}
	nilNodeInt := &node[int]{
		color: black,
	}

	testInt := testCase[int]{
		name: "int empty tree",
		want: &Tree[int]{
			root:    nilNodeInt,
			nilNode: nilNodeInt,
		},
	}
	t.Run(testInt.name, func(t *testing.T) {
		if got := New[int](); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})

	nilNodeStr := &node[string]{
		color: black,
	}
	testString := testCase[string]{
		name: "string empty tree",
		want: &Tree[string]{
			root:    nilNodeStr,
			nilNode: nilNodeStr,
		},
	}
	t.Run(testString.name, func(t *testing.T) {
		if got := New[int](); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})
}

func TestTree_NewWithElement(t *testing.T) {
	type args[V constraints.Ordered] struct {
		key   V
		value any
	}
	type testCase[V constraints.Ordered] struct {
		name string
		args args[V]
		want *Tree[V]
	}

	nilNodeInt := &node[int]{
		color: black,
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
					color:  black,
					left:   nilNodeInt,
					right:  nilNodeInt,
					parent: nilNodeInt,
				},
				nilNode: nilNodeInt,
			},
		},
		{
			name: "one element",
			args: args[int]{key: 15, value: 15},
			want: &Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
					color:  black,
					left:   nilNodeInt,
					right:  nilNodeInt,
					parent: nilNodeInt,
				},
				nilNode: nilNodeInt,
			},
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
		t    *Tree[V]
		want V
	}

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    getTree([]int{}),
			want: 0,
		},
		{
			name: "tree with one element",
			t:    getTree([]int{15}),
			want: 15,
		},
		{
			name: "tree with root and one element",
			t:    getTree([]int{15, 25}),
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
		t    *Tree[V]
		want V
	}

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    getTree([]int{}),
			want: 0,
		},
		{
			name: "tree with one element",
			t:    getTree([]int{15}),
			want: 15,
		},
		{
			name: "tree with root and one element",
			t:    getTree([]int{15, 25}),
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
		t    *Tree[V]
		args args[V]
		want bool
	}

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    getTree([]int{}),
			args: args[int]{key: 1},
			want: false,
		},
		{
			name: "tree with one element - not found",
			t:    getTree([]int{15}),
			args: args[int]{key: 1},
			want: false,
		},
		{
			name: "tree with one element - found",
			t:    getTree([]int{15}),
			args: args[int]{key: 15},
			want: true,
		},
		{
			name: "tree with root and one element - found",
			t:    getTree([]int{15, 25}),
			args: args[int]{key: 25},
			want: true,
		},
		{
			name: "tree with root and one element - not found",
			t:    getTree([]int{15, 25}),
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
		t       *Tree[V]
		args    args[V]
		want    any
		wantErr bool
	}

	tests := []testCase[int]{
		{
			name:    "empty tree",
			t:       getTree([]int{}),
			args:    args[int]{key: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "tree with one element - not found",
			t:       getTree([]int{}),
			args:    args[int]{key: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "tree with one element - found",
			t:       getTree([]int{15}),
			args:    args[int]{key: 15},
			want:    15,
			wantErr: false,
		},
		{
			name:    "tree with root and one element - found",
			t:       getTree([]int{15, 25}),
			args:    args[int]{key: 25},
			want:    25,
			wantErr: false,
		},
		{
			name:    "tree with root and one element - not found",
			t:       getTree([]int{15, 25}),
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

	tree := New[int]()

	treeWithOneRightElement := Tree[int]{
		root: &node[int]{
			element: element[int]{
				key:   15,
				value: 15,
			},
			parent: tree.nilNode,
			left:   tree.nilNode,
			right: &node[int]{
				element: element[int]{
					key:   25,
					value: 25,
				},
				parent: tree.nilNode,
				color:  red,
				left:   tree.nilNode,
				right:  tree.nilNode,
			},
			color: black,
		},
		nilNode: tree.nilNode,
	}
	treeWithOneRightElement.root.right.parent = treeWithOneRightElement.root

	treeWithOneLeftElement := Tree[int]{
		root: &node[int]{
			element: element[int]{
				key:   15,
				value: 15,
			},
			parent: tree.nilNode,
			left: &node[int]{
				element: element[int]{
					key:   10,
					value: 10,
				},
				parent: tree.nilNode,
				color:  red,
				left:   tree.nilNode,
				right:  tree.nilNode,
			},
			right: tree.nilNode,
			color: black,
		},
		nilNode: tree.nilNode,
	}
	treeWithOneLeftElement.root.left.parent = treeWithOneLeftElement.root

	tests := []testCase[int]{
		{
			name: "case 1: insert root",
			t:    tree,
			args: args[int]{key: 15, value: 15},
			want: &Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
					color:  black,
					left:   tree.nilNode,
					right:  tree.nilNode,
					parent: tree.nilNode,
				},
				nilNode: tree.nilNode,
			},
		},
		{
			name: "case 2 - insert right node to black root",
			t:    getTree([]int{15}),
			args: args[int]{key: 25, value: 25},
			want: &treeWithOneRightElement,
		},
		{
			name: "case 2 - insert left node to black root",
			t:    getTree([]int{15}),
			args: args[int]{key: 10, value: 10},
			want: &treeWithOneLeftElement,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Insert(tt.args.key, tt.args.value)
			if !treeEquals(tt.t, tt.want) {
				t1.Errorf("Insert() = %#+v, want %#+v", tt.t, tt.want)
			}
		})
	}
}

func TestTree_Insert_case_2(t1 *testing.T) {
	t := getTree([]int{11, 9, 18, 8, 10})

	// check tree's structure and colours before insert
	validTree := []validNode[int]{
		{node: t.root, key: 11, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 9, color: black, nodePath: "t.root.left"},
		{node: t.root.right, key: 18, color: black, nodePath: "t.root.right"},
		{node: t.root.left.left, key: 8, color: red, nodePath: "t.root.left.left"},
		{node: t.root.left.right, key: 10, color: red, nodePath: "t.root.left.right"},
	}
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// add list
	t.Insert(7, 7)

	// check tree's structure and colours after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 11, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 9, color: red, nodePath: "t.root.left"},
		{node: t.root.right, key: 18, color: black, nodePath: "t.root.right"},
		{node: t.root.left.left, key: 8, color: black, nodePath: "t.root.left.left"},
		{node: t.root.left.right, key: 10, color: black, nodePath: "t.root.left.right"},
		{node: t.root.left.left.left, key: 7, color: red, nodePath: "t.root.left.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Insert_case_3(t1 *testing.T) {
	t := getTree([]int{5, 3, 6})

	// check tree's structure and colours before insert
	validTree := []validNode[int]{
		{node: t.root, key: 5, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 3, color: red, nodePath: "t.root.left"},
		{node: t.root.right, key: 6, color: red, nodePath: "t.root.right"},
	}

	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// add list
	t.Insert(2, 2)

	// check tree's structure and colours after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 5, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 3, color: black, nodePath: "t.root.left"},
		{node: t.root.right, key: 6, color: black, nodePath: "t.root.right"},
		{node: t.root.left.left, key: 2, color: red, nodePath: "t.root.left.left"},
	}

	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Insert_case_4_left_rotate(t1 *testing.T) {
	t := getTree([]int{8, 6})

	validTree := []validNode[int]{
		{node: t.root, key: 8, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 6, color: red, nodePath: "t.root.left"},
	}

	// check tree's structure and colours before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// add list
	t.Insert(7, 7)

	// check tree's structure and colours after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 7, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 6, color: red, nodePath: "t.root.left"},
		{node: t.root.right, key: 8, color: red, nodePath: "t.root.right"},
	}

	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Insert_case_5_right_rotate(t1 *testing.T) {
	t := getTree([]int{8, 7})

	validTree := []validNode[int]{
		{node: t.root, key: 8, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 7, color: red, nodePath: "t.root.left"},
	}

	// check tree's structure and colours before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// add list
	t.Insert(6, 6)

	// check tree's structure and colours after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 7, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 6, color: red, nodePath: "t.root.left"},
		{node: t.root.right, key: 8, color: red, nodePath: "t.root.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Insert_big_case(t1 *testing.T) {
	t := getTree([]int{11, 2, 14, 1, 7, 15, 5, 8})

	validTree := []validNode[int]{
		{node: t.root, key: 11, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 2, color: red, nodePath: "t.root.left"},
		{node: t.root.right, key: 14, color: black, nodePath: "t.root.right"},
		{node: t.root.right.right, key: 15, color: red, nodePath: "t.root.right.right"},
		{node: t.root.left.left, key: 1, color: black, nodePath: "t.root.left.left"},
		{node: t.root.left.right, key: 7, color: black, nodePath: "t.root.left.right"},
		{node: t.root.left.right.left, key: 5, color: red, nodePath: "t.root.left.right.left"},
		{node: t.root.left.right.right, key: 8, color: red, nodePath: "t.root.left.right.right"},
	}

	// check tree's structure and colours before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// add list
	t.Insert(4, 4)

	// check tree's structure and colours after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 7, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 2, color: red, nodePath: "t.root.left"},
		{node: t.root.right, key: 11, color: red, nodePath: "t.root.right"},
		{node: t.root.left.left, key: 1, color: black, nodePath: "t.root.left.left"},
		{node: t.root.left.right, key: 5, color: black, nodePath: "t.root.left.right"},
		{node: t.root.right.left, key: 8, color: black, nodePath: "t.root.right.left"},
		{node: t.root.right.right, key: 14, color: black, nodePath: "t.root.right.right"},
		{node: t.root.right.right.right, key: 15, color: red, nodePath: "t.root.right.right.right"},
		{node: t.root.left.right.left, key: 4, color: red, nodePath: "t.root.left.right.left"},
	}

	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

//func TestTree_Insert_Cases(t1 *testing.T) {
//	type args[V constraints.Ordered] struct {
//		key   V
//		value V
//	}
//	type testCase[V constraints.Ordered] struct {
//		name                 string
//		t                    *Tree[V]
//		args                 args[V]
//		validTree            []validNode[int]
//		validTreeAfterInsert []validNode[int]
//	}
//
//	treeCase2 := getTree([]int{11, 9, 18, 8, 10})
//	tests := []testCase[int]{
//		{
//			name: "case 2",
//			t:    treeCase2,
//			args: args[int]{key: 7, value: 7},
//			validTree: []validNode[int]{
//				{node: treeCase2.root, key: 11, color: black, nodePath: "t.root"},
//				{node: treeCase2.root.left, key: 9, color: black, nodePath: "t.root.left"},
//				{node: treeCase2.root.right, key: 18, color: black, nodePath: "t.root.right"},
//				{node: treeCase2.root.left.left, key: 8, color: red, nodePath: "t.root.left.left"},
//				{node: treeCase2.root.left.right, key: 10, color: red, nodePath: "t.root.left.right"},
//			},
//			validTreeAfterInsert: []validNode[int]{
//				{node: treeCase2.root, key: 11, color: black, nodePath: "t.root"},
//				{node: treeCase2.root.left, key: 9, color: red, nodePath: "t.root.left"},
//				{node: treeCase2.root.right, key: 18, color: black, nodePath: "t.root.right"},
//				{node: treeCase2.root.left.left, key: 8, color: black, nodePath: "t.root.left.left"},
//				{node: treeCase2.root.left.right, key: 10, color: black, nodePath: "t.root.left.right"},
//				{node: treeCase2.root.left.left.left, key: 7, color: red, nodePath: "t.root.left.right"},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t1.Run(tt.name, func(t1 *testing.T) {
//			for _, n := range tt.validTree {
//				checkNode(t1, &n)
//			}
//			tt.t.Insert(tt.args.key, tt.args.value)
//			for _, n := range tt.validTreeAfterInsert {
//				checkNode(t1, &n)
//			}
//		})
//	}
//}

func TestTree_Delete(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    *Tree[V]
		args args[V]
		want *Tree[V]
	}

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    getTree([]int{}),
			args: args[int]{key: 1},
			want: getTree([]int{}),
		},
		{
			name: "tree only with root - without changes",
			t:    getTree([]int{15}),
			args: args[int]{key: 1},
			want: getTree([]int{15}),
		},
		{
			name: "tree only with root - delete root",
			t:    getTree([]int{15}),
			args: args[int]{key: 15},
			want: getTree([]int{}),
		},
		{
			name: "tree with elements - without changes",
			t:    getTree([]int{15, 25}),
			args: args[int]{key: 85},
			want: getTree([]int{15, 25}),
		},
		{
			name: "tree with elements - delete node without children",
			t:    getTree([]int{15, 25}),
			args: args[int]{key: 25},
			want: getTree([]int{15}),
		},
		{
			name: "delete root with left and right node",
			t:    getTree([]int{25, 15, 35}),
			args: args[int]{key: 25},
			want: getTree([]int{35, 15}),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Delete(tt.args.key)
			if !treeEquals(tt.t, tt.want) {
				t1.Errorf("Delete() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func TestTree_Delete_case_1_delete_red_list(t1 *testing.T) {
	t := getTree([]int{10, 7, 11, 8})

	validTree := []validNode[int]{
		{node: t.root, key: 10, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 7, color: black, nodePath: "t.root.left"},
		{node: t.root.right, key: 11, color: black, nodePath: "t.root.right"},
		{node: t.root.left.right, key: 8, color: red, nodePath: "t.root.left.right"},
	}

	// check tree's structure and colours before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// delete list
	t.Delete(8)

	// check tree's structure and colours after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 10, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 7, color: black, nodePath: "t.root.left"},
		{node: t.root.right, key: 11, color: black, nodePath: "t.root.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
	checkNodeIsNilNode(t1, t.nilNode, t.root.left.right, "t.root.left.right")
}

func TestTree_Delete_case_3_delete_black_node_with_one_red_child(t1 *testing.T) {
	t := getTree([]int{10, 7, 11, 8})

	validTree := []validNode[int]{
		{node: t.root, key: 10, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 7, color: black, nodePath: "t.root.left"},
		{node: t.root.right, key: 11, color: black, nodePath: "t.root.right"},
		{node: t.root.left.right, key: 8, color: red, nodePath: "t.root.left.right"},
	}

	// check tree's structure and colours before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// delete list
	t.Delete(7)

	// check tree's structure and colours after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 10, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 8, color: black, nodePath: "t.root.left"},
		{node: t.root.right, key: 11, color: black, nodePath: "t.root.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Delete_case_2(t1 *testing.T) {
	t := getTree([]int{10, 8, 12, 11, 14})

	validTree := []validNode[int]{
		{node: t.root, key: 10, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 8, color: black, nodePath: "t.root.left"},
		{node: t.root.right, key: 12, color: black, nodePath: "t.root.right"},
		{node: t.root.right.left, key: 11, color: red, nodePath: "t.root.right.left"},
		{node: t.root.right.right, key: 14, color: red, nodePath: "t.root.right.right"},
	}

	// check tree's structure and colours before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// delete list
	t.Delete(8)

	// check tree's structure and colours after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 12, color: black, nodePath: "t.root"},
		{node: t.root.left, key: 10, color: black, nodePath: "t.root.left"},
		{node: t.root.left.right, key: 11, color: red, nodePath: "t.root.left.right"},
		{node: t.root.right, key: 14, color: black, nodePath: "t.root.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func checkNode[V constraints.Ordered](t *testing.T, vn *validNode[V]) {
	if vn == nil {
		return
	}

	if vn.node.element.key != vn.key {
		t.Errorf("Error - Want key: %v, have key %v in %s\n",
			vn.key,
			vn.node.element.key,
			vn.nodePath,
		)
	}

	if vn.node.color != vn.color {
		t.Errorf("Error - Want color: %v, have color: %v in %s\n",
			vn.color,
			vn.node.color,
			vn.nodePath,
		)
	}
}

func checkNodeIsNilNode(t *testing.T, nilNode, node *node[int], nodePath string) {
	if node != nilNode {
		t.Errorf("Error - Node is not nil node in %s\n", nodePath)
	}
}

func getTree(elements []int) *Tree[int] {
	tree := New[int]()
	for _, el := range elements {
		tree.Insert(el, el)
	}

	return tree
}

func treeEquals(tree1, tree2 *Tree[int]) bool {
	return nodesEquals(tree1.root, tree2.root)
}

func nodesEquals(node1, node2 *node[int]) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	return node1.element.key == node2.element.key &&
		node1.element.value == node2.element.value &&
		node1.color == node2.color &&
		nodesEquals(node1.left, node2.left) &&
		nodesEquals(node1.right, node2.right)
}
