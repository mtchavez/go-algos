package rbt

type Color int

const (
	RED Color = iota
	BLACK
	DOUBLY_BLACK
)

var COLORMAP = map[Color]string{
	RED:          "RED",
	BLACK:        "BLACK",
	DOUBLY_BLACK: "DOUBLY_BLACK",
}

type node struct {
	color  Color
	left   *node
	right  *node
	parent *node
	key    int
}

func (n *node) setLeft(x *node) {
	n.left = x
	if x != nil {
		x.parent = n
	}
}

func (n *node) setRight(x *node) {
	n.right = x
	if x != nil {
		x.parent = n
	}
}

func (n *node) setChildren(x, y *node) {
	n.setLeft(x)
	n.setRight(y)
}

func (n *node) replace(y *node) {
	if y.parent == nil {
		if y != nil {
			y.parent = nil
		}
	} else if n.parent.left == n {
		n.parent.setLeft(y)
	} else {
		n.parent.setRight(y)
	}
	n.parent = nil
}

func (n *node) sibling() *node {
	if n.parent.left == n {
		return n.parent.right
	} else {
		return n.parent.left
	}
}

func (n *node) uncle() *node {
	if n.parent == nil {
		return n.parent
	}
	return n.parent.sibling()
}

func (n *node) grandparent() *node {
	if n.parent == nil {
		return n.parent
	}
	return n.parent.parent
}
