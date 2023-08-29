rbtree
=======================

Library for work with Red-black trees.

You can create a Red-black tree and use a list of functions to work with it.

## Tree functions
- [Empty tree's creation example](#empty-trees-creation-example)
- [Tree's creation with one element example](#trees-creation-with-one-element-example)
- [Insert element to tree](#insert-element-to-tree)
- [Exists element](#exists-element)
- [Get value by key element](#get-value-by-key-element)
- [Min tree element](#min-tree-element)
- [Max tree element](#max-tree-element)


### Empty tree's creation example

```
t := tree.New[int]() // empty int tree
t := tree.New[string]() // empty string tree
```

### Tree's creation with one element example

```
t := tree.NewWithElement[int](1,1) // int tree creation with one element
t := tree.NewWithElement[string]("key", "value") // string tree creation with one element
```

### Insert element to tree
```
t := tree.New[int]() // empty int tree
t.Insert(22, 22) // insert to tree element: key=22, value=22
t.Insert(8, 8) // insert to tree element: key=8, value=8
t.Insert(4, 4) // insert to tree element: key=4, value=4
```

### Exists element

```
t := tree.New[int]()
t.Insert(22, 22)
t.Insert(8, 8)
t.Insert(4, 4)

resultNil := t.Exists(15) // false
result    := t.Exists(8)  // true
```

### Get value by key element

```
t := tree.New[int]()
t.Insert(22, 22)
t.Insert(8, 8)
t.Insert(4, 4)

resultNil, err := t.GetValue(15) // nil, err
result, err    := t.GetValue(8)  // 8, nil
```

### Min tree element
```
t := tree.New[int]()
t.Insert(22, 22)
t.Insert(8, 8)
t.Insert(4, 4)

result := t.Min() // 4
```
### Max tree element
```
t := tree.New[int]()
t.Insert(22, 22)
t.Insert(8, 8)
t.Insert(4, 4)

result := t.Max() // 22
```