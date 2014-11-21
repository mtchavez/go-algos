package rbt

import "fmt"

// Tree is a struct that references
// the root node of the RB Tree
type Tree struct {
	root *node
}

func (t *Tree) leftRotate(x *node) {
	parent, y := x.parent, x.right
	var a, b, c *node
	if y != nil {
		b = y.left
		c = y.right
		x.replace(y)
		y.setChildren(y, c)
	}
	a = x.left
	x.setChildren(a, b)
	if parent == nil && y != nil {
		t.root = y
	}
}

func (t *Tree) rightRotate(y *node) {
	parent, x := y.parent, y.left
	var a, b, c *node
	if x != nil {
		a = x.left
		b = x.right
		y.replace(x)
		x.setChildren(a, y)
	}
	c = y.right
	y.setChildren(b, c)
	if parent == nil && x != nil {
		t.root = x
	}
}

// Insert takes key to put into tree
func (t *Tree) Insert(key int) {
	n := t.root
	x := &node{key: key}
	var parent *node
	for n != nil {
		parent = n
		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	if parent == nil {
		// Tree is empty so set root
		t.root = x
	} else if key < parent.key {
		parent.setLeft(x)
	} else {
		parent.setRight(x)
	}
	t.balanceInsert(x)
}

func (t *Tree) balanceInsert(x *node) {
	for x.parent != nil && x.parent.color == Red {
		uncle := x.uncle()
		grandparent := x.grandparent()
		if uncle != nil && uncle.color == Red {
			x.parent.color = Black
			grandparent.color = Red
			uncle.color = Black
			x = grandparent
		} else {
			if x.parent == grandparent.left {
				if x == x.parent.right {
					x = x.parent
					t.leftRotate(x)
				}
				x.parent.color = Black
				grandparent.color = Red
				t.rightRotate(grandparent)
			} else {
				if x == x.parent.left {
					x = x.parent
					t.rightRotate(x)
				}
				x.parent.color = Black
				grandparent.color = Red
				t.rightRotate(grandparent)
			}
		}
	}
	t.root.color = Black
}

func (t *Tree) removeNode(x *node) {
	if x == nil {
		return
	}
	x.parent, x.left, x.right = nil, nil, nil
}

// Min returns min value of tree
func (t *Tree) Min() *node {
	r := t.root
	for r != nil && r.left != nil {
		r = r.left
	}
	return r
}

// Search takes a key to search the tree for
// and will return a node if found or nil
func (t *Tree) Search(key int) *node {
	r := t.root
	for r != nil && r.key != key {
		if key < r.key {
			r = r.left
		} else {
			r = r.right
		}
	}
	return r
}

func (t *Tree) String() string {
	return t.asString(t.root)
}

func (t *Tree) asString(n *node) string {
	if n == nil {
		return "."
	}
	return fmt.Sprintf("(l:%+v key:%+v color:%+v r:%+v)", t.asString(n.left), n.key, ColorMap[n.color], t.asString(n.right))
}
