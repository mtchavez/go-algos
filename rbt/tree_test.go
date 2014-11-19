package rbt

import "testing"

func TestMin_empty(t *testing.T) {
	tree := &Tree{}
	min := tree.Min()
	if min != nil {
		t.Errorf("Min of empty tree should be nil")
	}
}

func TestMin(t *testing.T) {
	tree := &Tree{}
	for _, val := range []int{5, 4, 3, 2, 1} {
		tree.Insert(val)
	}
	min := tree.Min()
	if min.key != 1 {
		t.Errorf("Min oftree should be 1 but got %+v", min.key)
	}
}

func TestSearch_empty(t *testing.T) {
	tree := &Tree{}
	found := tree.Search(99)
	if found != nil {
		t.Errorf("Should not have found a node for an empty tree")
	}
}

func TestSearch_notfound(t *testing.T) {
	tree := &Tree{}
	for _, val := range []int{1, 2, 3, 4, 5} {
		tree.Insert(val)
	}
	found := tree.Search(99)
	if found != nil {
		t.Errorf("Should not have found a node that is not in the tree")
	}
}

func TestSearch_found(t *testing.T) {
	tree := &Tree{}
	for _, val := range []int{1, 2, 3, 4, 5} {
		tree.Insert(val)
	}
	found := tree.Search(4)
	if found == nil {
		t.Errorf("Should have found a node that is in the tree")
	}
	if found.key != 4 {
		t.Errorf("Found key should have been %v but got %v", 4, found.key)
	}
}

func TestSearch_foundLessKey(t *testing.T) {
	tree := &Tree{}
	for _, val := range []int{5, 4, 3, 2, 1} {
		tree.Insert(val)
	}
	found := tree.Search(2)
	if found == nil {
		t.Errorf("Should have found a node that is in the tree")
	}
	if found.key != 2 {
		t.Errorf("Found key should have been %v but got %v", 2, found.key)
	}
}

func TestRemoveNode(t *testing.T) {
	tree := &Tree{}
	for _, val := range []int{5, 4, 3, 2, 1} {
		tree.Insert(val)
	}
	n := tree.Search(2)
	tree.removeNode(n)
	if n.parent != nil || n.left != nil || n.right != nil {
		t.Errorf("Removed node should have linked nodes set to nil")
	}
}
