package rbt

import "testing"

func Test_setLeft(t *testing.T) {
	n := &node{}
	left := &node{}
	n.setLeft(left)
	if left.parent != n {
		t.Errorf("Expected parent of left node to be set")
	}
}

func Test_setLeft_nil(t *testing.T) {
	n := &node{}
	n.setLeft(nil)
	if n.left != nil {
		t.Errorf("Expected left node to be nil")
	}
}

func Test_setRight(t *testing.T) {
	n := &node{}
	right := &node{}
	n.setRight(right)
	if right.parent != n {
		t.Errorf("Expected parent of right node to be set")
	}
}

func Test_setRight_nil(t *testing.T) {
	n := &node{}
	n.setRight(nil)
	if n.right != nil {
		t.Errorf("Expected right node to be nil")
	}
}

func Test_setChildren(t *testing.T) {
	n := &node{}
	left := &node{}
	right := &node{}
	n.setChildren(left, right)

	if n.left != left {
		t.Errorf("Expected left node to be set")
	}
	if n.right != right {
		t.Errorf("Expected right node to be set")
	}

	if right.parent != n {
		t.Errorf("Expected right parent to be set")
	}
	if left.parent != n {
		t.Errorf("Expected left parent to be set")
	}
}

func Test_replace_noParent(t *testing.T) {
	n := &node{}
	n.replace(nil)
	if n.parent != nil {
		t.Errorf("Expected parent to be nil after replace")
	}

	p := &node{}
	y := &node{parent: p}
	n.replace(y)
	if y.parent != nil {
		t.Errorf("Expected y node to have no parent")
	}
	if n.parent != nil {
		t.Errorf("Parent of n should be nil")
	}
}

func Test_replace_parentLeft(t *testing.T) {
	p := &node{}
	n := &node{}
	p.setChildren(n, nil)
	y := &node{}
	n.replace(y)
	if n.parent != nil {
		t.Errorf("Parent of n should be nil on replace")
	}
	if p.left != y {
		t.Errorf("Left parent node should have been replaced with y")
	}
}

func Test_replace_parentRight(t *testing.T) {
	p := &node{}
	n := &node{}
	p.setChildren(nil, n)
	y := &node{}
	n.replace(y)
	if n.parent != nil {
		t.Errorf("Parent of n should be nil on replace")
	}
	if p.right != y {
		t.Errorf("Right arent node should have been replaced with y")
	}
}

func Test_sibling_noParent(t *testing.T) {
	n := &node{}
	if n.sibling() != nil {
		t.Errorf("Sibling of node without parent should be nil")
	}
}

func Test_sibling_leftNodeSame(t *testing.T) {
	x := &node{}
	y := &node{}
	y.setLeft(x)
	if x.sibling() != nil {
		t.Errorf("Sibling of left node of parent should be nil if no right node")
	}

	right := &node{}
	y.setRight(right)
	if x.sibling() != right {
		t.Errorf("Sibling of left node should be right node")
	}
}

func Test_sibling_leftNode(t *testing.T) {
	x := &node{}
	y := &node{}
	z := &node{}
	y.setChildren(z, x)
	if x.sibling() != z {
		t.Errorf("Sibling of left node of parent should be nil if no right node")
	}
}

func Test_uncle_nilParent(t *testing.T) {
	x := &node{}
	if x.uncle() != nil {
		t.Errorf("Uncle of node without parent should be nil")
	}
}

func Test_grandparent(t *testing.T) {
	x := &node{}
	n := &node{}
	x.setChildren(n, nil)
	if x.grandparent() != nil {
		t.Errorf("Grandparent of node without parent should be nil")
	}
	y := &node{}
	n.setChildren(y, nil)
	if y.grandparent() != x {
		t.Errorf("Granparent of node should be the parent of the parent")
	}
}
