package algorithms_test

import (
	"fmt"

	"git.sr.ht/~telemachus/algorithms"
)

func makeHeap() *algorithms.Heap {
	h := algorithms.NewHeap()
	h.Insert(15)
	h.Insert(6)
	h.Insert(18)
	h.Insert(3)
	h.Insert(2)
	h.Insert(4)
	h.Insert(7)
	h.Insert(13)
	h.Insert(9)
	h.Insert(17)
	h.Insert(20)
	h.Insert(7)
	return h
}

func ExampleLoopThroughHeap() {
	h := makeHeap()

	for n, ok := h.Delete(); ok != false; n, ok = h.Delete() {
		fmt.Println(n)
	}
	// Output:
	// 20
	// 18
	// 17
	// 15
	// 13
	// 9
	// 7
	// 7
	// 6
	// 4
	// 3
	// 2
}
