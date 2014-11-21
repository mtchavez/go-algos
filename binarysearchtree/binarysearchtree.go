package bst

type node struct {
	key    int
	parent *node
	left   *node
	right  *node
}

// Tree is a binary search tree
// containing root node of tree
type Tree struct {
	root *node
}

// ValConsumer function to take val
// from walking tree functions
type ValConsumer func(val int)

// New creates a new binary search tree
func New() *Tree {
	return &Tree{}
}

// InOrderWalk will walk through the tree
// and takes a ValConsumer function to handle
// the value of each node as they are walked
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

// Search for a value in the tree and returns
// the node if found
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

// Min returns min value in tree
func (t *Tree) Min() *node {
	return min(t.root)
}

func min(n *node) *node {
	for n != nil && n.left != nil {
		n = n.left
	}
	return n
}

// Max returns max value in tree
func (t *Tree) Max() (n *node) {
	return max(t.root)
}

func max(n *node) *node {
	for n != nil && n.right != nil {
		n = n.right
	}
	return n
}

// Successor of x returns the smallest value y
// from element x such that y > x
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

// Predecessor of x returns the largest value y
// from element x such that y < x
func Predecessor(n *node) (s *node) {
	if n == nil {
		return s
	}
	if n.left != nil {
		return min(n.left)
	}
	s = n.parent
	for s != nil && s.right != n {
		n = s
		s = s.parent
	}
	return s
}

// Insert takes an integer value
// to insert into the tree
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
