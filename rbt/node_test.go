package rbt

import (
	"testing"
)

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
