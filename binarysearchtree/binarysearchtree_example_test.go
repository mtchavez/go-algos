package bst

import "fmt"

func Example_InOrderWalk() {
	t := New()
	values := []int{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}
	for _, x := range values {
		t.Insert(x)
	}
	consumed := []int{}
	p := ValConsumer(func(val int) {
		consumed = append(consumed, val)
	})
	t.InOrderWalk(p)
	fmt.Printf("Consumed %+v\n", consumed)
	// Output:
	// Consumed [1 2 3 4 7 8 9 10 14 16]
}

func Example_Search() {
	t := New()
	values := []int{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}
	for _, x := range values {
		t.Insert(x)
	}
	node := t.Search(16)
	fmt.Printf("Searched for 16 and got %+v\n", node.key)
	// Output:
	// Searched for 16 and got 16
}

func Example_Min() {
	t := New()
	values := []int{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}
	for _, x := range values {
		t.Insert(x)
	}
	node := t.Min()
	fmt.Printf("Min of tree is %+v\n", node.key)
	// Output:
	// Min of tree is 1
}

func Example_Max() {
	t := New()
	values := []int{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}
	for _, x := range values {
		t.Insert(x)
	}
	node := t.Max()
	fmt.Printf("Max of tree is %+v\n", node.key)
	// Output:
	// Max of tree is 16
}

func Example_Successor() {
	t := New()
	values := []int{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}
	for _, x := range values {
		t.Insert(x)
	}
	node := Successor(t.Search(8))
	fmt.Printf("Successor of 8 is %+v\n", node.key)
	// Output:
	// Successor of 8 is 9
}
