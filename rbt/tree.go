package rbt

import "fmt"

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
	for x.parent != nil && x.parent.color == RED {
		uncle := x.uncle()
		grandparent := x.grandparent()
		if uncle != nil && uncle.color == RED {
			x.parent.color = BLACK
			grandparent.color = RED
			uncle.color = BLACK
			x = grandparent
		} else {
			if x.parent == grandparent.left {
				if x == x.parent.right {
					x = x.parent
					t.leftRotate(x)
				}
				x.parent.color = BLACK
				grandparent.color = RED
				t.rightRotate(grandparent)
			} else {
				if x == x.parent.left {
					x = x.parent
					t.rightRotate(x)
				}
				x.parent.color = BLACK
				grandparent.color = RED
				t.rightRotate(grandparent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *Tree) removeNode(x *node) {
	if x == nil {
		return
	}
	x.parent, x.left, x.right = nil, nil, nil
}

func (t *Tree) Min() *node {
	r := t.root
	for r != nil && r.left != nil {
		r = r.left
	}
	return r
}

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
	} else {
		return fmt.Sprintf("(l:%+v key:%+v color:%+v r:%+v)", t.asString(n.left), n.key, COLORMAP[n.color], t.asString(n.right))
	}

}
