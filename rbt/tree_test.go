package rbt

import "testing"

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
