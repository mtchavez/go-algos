package avl

type node struct {
	delta  int
	key    int
	left   *node
	right  *node
	parent *node
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
