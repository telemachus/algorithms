package algorithms_test

import (
	"fmt"

	"github.com/telemachus/algorithms"
	"golang.org/x/tour/tree"
)

func ExampleTree_RecursiveConcurrentWalk() {
	ch := make(chan int)
	go algorithms.RecursiveConcurrentWalk(tree.New(1), ch)

	for n := range ch {
		fmt.Println(n)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}

func ExampleTree_IterativeConcurrentWalk() {
	ch := make(chan int)
	go algorithms.IterativeConcurrentWalk(tree.New(2), ch)

	for n := range ch {
		fmt.Println(n)
	}
	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
	// 12
	// 14
	// 16
	// 18
	// 20
}

func ExampleTree_MorrisIterativeConcurrentWalk() {
	ch := make(chan int)
	go algorithms.MorrisIterativeConcurrentWalk(tree.New(2), ch)

	for n := range ch {
		fmt.Println(n)
	}
	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
	// 12
	// 14
	// 16
	// 18
	// 20
}
