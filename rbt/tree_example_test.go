package rbt

import "fmt"

func ExampleTree_String() {
	values := []int{13, 8, 17, 1, 11, 15, 25, 6, 22, 27}
	tree := Tree{}
	for _, v := range values {
		tree.Insert(v)
	}
	fmt.Printf(tree.String())
	// Output:
	// (l:(l:(l:. key:1 color:BLACK r:(l:. key:6 color:RED r:.)) key:8 color:RED r:(l:. key:11 color:BLACK r:.)) key:13 color:BLACK r:(l:(l:. key:15 color:BLACK r:.) key:17 color:RED r:(l:(l:. key:22 color:RED r:.) key:25 color:BLACK r:(l:. key:27 color:RED r:.))))
}

func ExampleTree_Delete() {
	values := []int{13, 8, 17, 1, 11, 15, 25, 6, 22, 27}
	tree := Tree{}
	for _, v := range values {
		tree.Insert(v)
	}
	n := tree.Search(15)
	tree.Delete(n)
	fmt.Printf("%+v\n", tree.String())

	n = tree.Search(6)
	tree.Delete(n)
	fmt.Printf("%+v\n", tree.String())
	// Output:
	// (l:(l:(l:. key:1 color:BLACK r:(l:. key:6 color:RED r:.)) key:8 color:RED r:(l:. key:11 color:BLACK r:.)) key:13 color:BLACK r:(l:. key:17 color:RED r:(l:(l:. key:22 color:RED r:.) key:25 color:BLACK r:(l:. key:27 color:RED r:.))))
	// (l:(l:(l:. key:1 color:BLACK r:.) key:8 color:RED r:(l:. key:11 color:BLACK r:.)) key:13 color:BLACK r:(l:. key:17 color:RED r:(l:(l:. key:22 color:RED r:.) key:25 color:BLACK r:(l:. key:27 color:RED r:.))))
}
