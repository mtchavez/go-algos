package rbt

type Tree struct {
	root *node
}

func (t *Tree) leftRotate(x *node) {
	parent, y := x.parent, x.right
	a, b, c := x.left, y.left, y.right
	x.replace(y)
	x.setChildren(a, b)
	y.setChildren(y, c)
	if parent == nil {
		t.root = y
	}
}

func (t *Tree) rightRotate(y *node) {
	parent, x := y.parent, y.left
	a, b, c := x.left, x.right, y.right
	y.replace(x)
	y.setChildren(b, c)
	x.setChildren(a, y)
	if parent == nil {
		t.root = x
	}
}
