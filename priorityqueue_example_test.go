package algorithms_test

import (
	"fmt"

	"git.sr.ht/~telemachus/algorithms"
)

func makePriorityQueue() *algorithms.PriorityQueue {
	items := make([]*algorithms.Item, 0, 10)
	items = append(items, &algorithms.Item{Name: "four", Priority: 4})
	items = append(items, &algorithms.Item{Name: "one", Priority: 1})
	items = append(items, &algorithms.Item{Name: "three", Priority: 3})
	items = append(items, &algorithms.Item{Name: "nine", Priority: 9})
	items = append(items, &algorithms.Item{Name: "five", Priority: 5})
	items = append(items, &algorithms.Item{Name: "two", Priority: 2})
	items = append(items, &algorithms.Item{Name: "sixteen", Priority: 16})
	items = append(items, &algorithms.Item{Name: "ten", Priority: 10})
	items = append(items, &algorithms.Item{Name: "fourteen", Priority: 14})
	items = append(items, &algorithms.Item{Name: "eight", Priority: 8})
	items = append(items, &algorithms.Item{Name: "seven", Priority: 7})

	pq := algorithms.NewPQ()
	pq.Init(items)

	return pq
}

func ExampleLoopOverPriorityQueue() {
	pq := makePriorityQueue()
	pq.Insert("nine", 9)
	if top, ok := pq.Max(); ok {
		fmt.Println(top.Priority)
	}

	for n, ok := pq.RemoveMax(); ok != false; n, ok = pq.RemoveMax() {
		fmt.Println(n.Priority)
	}
	// Output:
	// 16
	// 16
	// 14
	// 10
	// 9
	// 9
	// 8
	// 7
	// 5
	// 4
	// 3
	// 2
	// 1
}
