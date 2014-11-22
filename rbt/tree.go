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

// Delete takes a node in the tree to be removed
func (t *Tree) Delete(x *node) {
	if x == nil {
		return
	}
	var parent, db *node
	if x.left == nil {
		x.replace(x.right)
		db = x.right
	} else if x.right == nil {
		x.replace(x.left)
		db = x.left
	} else {
		y := x.min()
		parent = y.parent
		db = y.right
		x.key = y.key
		y.replace(y.right)
		x = y
	}
	if x.color == Black {
		bal := db
		if db == nil && parent != nil {
			if parent.isLeaf() {
				parent.color = DoubleBlack
			}
			bal = parent
		} else if db != nil {
			db.color = db.color + 1
		}
		t.balanceDelete(bal)
	}
	t.removeNode(x)
}

func checkNodeColor(n *node, c Color) bool {
	if n == nil {
		return false
	}
	return n.color == c
}

func (t *Tree) balanceDelete(db *node) {
	x := t.root
	if db == nil {
		return
	}
	for db != x && db.color == DoubleBlack {
		sibling := db.sibling()
		if sibling != nil {
			if checkNodeColor(sibling, Red) {
				db.parent.color = Red
				sibling.color = Black
				if db == db.parent.left {
					t.leftRotate(db.parent)
				} else {
					t.rightRotate(db.parent)
				}
			} else if checkNodeColor(sibling, Black) && checkNodeColor(sibling.left, Red) {
				prevColor := db.parent.color
				if db == db.parent.left {
					db.color = Black
					db.parent.color = Black
					sibling.left.color = prevColor
					t.rightRotate(sibling)
					t.leftRotate(db.parent)
				} else {
					db.color = Black
					db.parent.color = Black
					sibling.color = prevColor
					sibling.left.color = Black
				}
			} else if checkNodeColor(sibling, Black) && checkNodeColor(sibling.right, Red) {
				prevColor := db.parent.color
				if db == db.parent.left {
					db.color = Black
					db.parent.color = Black
					sibling.color = prevColor
					sibling.right.color = Black
					t.leftRotate(db.parent)
				} else {
					db.color = Black
					db.parent.color = Black
					sibling.right.color = prevColor
					t.leftRotate(sibling)
					t.rightRotate(db.parent)
				}
			} else if checkNodeColor(sibling, Black) && !checkNodeColor(sibling.left, Red) && !checkNodeColor(sibling.right, Red) {
				db.color = Black
				sibling.color = Red
				db.parent.color = db.parent.color + 1
				db = db.parent
			}
		} else {
			db.color = Black
			db.parent.color = db.parent.color + 1
			db = db.parent
		}
	}

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
