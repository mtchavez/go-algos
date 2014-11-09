package bst

type node struct {
	key    int
	parent *node
	left   *node
	right  *node
}

type Tree struct {
	root *node
}

type ValConsumer func(val int)

func New() *Tree {
	return &Tree{}
}

func (t *Tree) InOrderWalk(f ValConsumer) {
	inOrder(t.root, f)
}

func inOrder(n *node, f ValConsumer) {
	if n == nil {
		return
	}
	inOrder(n.left, f)
	f(n.key)
	inOrder(n.right, f)
}

func (t *Tree) Search(val int) (n *node) {
	n = t.root
	for n != nil && n.key != val {
		if val < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}

func (t *Tree) Min() *node {
	return min(t.root)
}

func min(n *node) *node {
	for n != nil && n.left != nil {
		n = n.left
	}
	return n
}

func (t *Tree) Max() (n *node) {
	n = t.root
	for n != nil && n.right != nil {
		n = n.right
	}
	return n
}

func Successor(n *node) (s *node) {
	if n == nil {
		return s
	}
	if n.right != nil {
		return min(n.right)
	}
	s = n.parent
	for s != nil && s.left != n {
		n = s
		s = s.parent
	}
	return s
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &node{key: val}
		return
	}
	insert(t.root, val)
}

func insert(n *node, val int) {
	newNode := &node{key: val}
	var parent *node
	for n != nil {
		parent = n
		if val < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	newNode.parent = parent
	if parent == nil {
		return
	} else if val < parent.key {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}
