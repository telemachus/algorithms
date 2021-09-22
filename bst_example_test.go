package algorithms_test

import (
	"fmt"

	"git.sr.ht/~telemachus/algorithms"
)

func ExampleInOrderWalk() {
	t := algorithms.NewBST()
	t.Insert(15)
	t.Insert(6)
	t.Insert(18)
	t.Insert(3)
	t.Insert(2)
	t.Insert(4)
	t.Insert(7)
	t.Insert(13)
	t.Insert(9)
	t.Insert(17)
	t.Insert(20)

	show := func(i int) {
		fmt.Printf("%d\n", i)
	}

	t.InOrderWalk(show)
	// Output:
	// 2
	// 3
	// 4
	// 6
	// 7
	// 9
	// 13
	// 15
	// 17
	// 18
	// 20
}

func ExamplePreOrderWalk() {
	t := algorithms.NewBST()
	t.Insert(15)
	t.Insert(6)
	t.Insert(18)
	t.Insert(3)
	t.Insert(2)
	t.Insert(4)
	t.Insert(7)
	t.Insert(13)
	t.Insert(9)
	t.Insert(17)
	t.Insert(20)

	show := func(i int) {
		fmt.Printf("%d\n", i)
	}

	t.PreOrderWalk(show)
	// Output:
	// 15
	// 6
	// 3
	// 2
	// 4
	// 7
	// 13
	// 9
	// 18
	// 17
	// 20
}

func ExamplePostOrderWalk() {
	t := algorithms.NewBST()
	t.Insert(15)
	t.Insert(6)
	t.Insert(18)
	t.Insert(3)
	t.Insert(2)
	t.Insert(4)
	t.Insert(7)
	t.Insert(13)
	t.Insert(9)
	t.Insert(17)
	t.Insert(20)

	show := func(i int) {
		fmt.Printf("%d\n", i)
	}

	t.PostOrderWalk(show)
	// Output:
	// 2
	// 4
	// 3
	// 9
	// 13
	// 7
	// 6
	// 17
	// 20
	// 18
	// 15
}
