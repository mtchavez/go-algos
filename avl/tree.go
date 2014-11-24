package avl

type Tree struct {
	root *node
}

func (t *Tree) Insert(key int) {
	x := &node{key: key}
	y := t.root
	var parent *node
	for y != nil {
		parent = y
		if key < y.key {
			y = y.left
		} else {
			y = y.right
		}
	}
	if parent == nil {
		t.root = x
	} else if key < parent.key {
		parent.setLeft(x)
	} else {
		parent.setRight(x)
	}
}
