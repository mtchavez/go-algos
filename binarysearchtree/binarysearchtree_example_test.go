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
