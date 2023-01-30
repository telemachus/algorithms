package algorithms_test

import (
	"fmt"

	"github.com/telemachus/algorithms"
)

func makeTree() *algorithms.Tree {
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
	t.Insert(7)
	return t
}

func show(i int) {
	fmt.Println(i)
}

func ExampleTree_InOrderWalk() {
	t := makeTree()
	t.InOrderWalk(show)
	// Output:
	// 2
	// 3
	// 4
	// 6
	// 7
	// 7
	// 9
	// 13
	// 15
	// 17
	// 18
	// 20
}

func ExampleTree_InOrderIterativeWalk() {
	t := makeTree()
	t.InOrderIterativeWalk(show)
	// Output:
	// 2
	// 3
	// 4
	// 6
	// 7
	// 7
	// 9
	// 13
	// 15
	// 17
	// 18
	// 20
}

func ExampleTree_PreOrderWalk() {
	t := makeTree()
	t.PreOrderWalk(show)
	// Output:
	// 15
	// 6
	// 3
	// 2
	// 4
	// 7
	// 7
	// 13
	// 9
	// 18
	// 17
	// 20
}

func ExampleTree_PreOrderIterativeWalk() {
	t := makeTree()
	t.PreOrderIterativeWalk(show)
	// Output:
	// 15
	// 6
	// 3
	// 2
	// 4
	// 7
	// 7
	// 13
	// 9
	// 18
	// 17
	// 20
}

func ExampleTree_PostOrderWalk() {
	t := makeTree()
	t.PostOrderWalk(show)
	// Output:
	// 2
	// 4
	// 3
	// 9
	// 13
	// 7
	// 7
	// 6
	// 17
	// 20
	// 18
	// 15
}

func ExampleTree_PostOrderIterativeWalk() {
	t := makeTree()
	t.PostOrderIterativeWalk(show)
	// Output:
	// 2
	// 4
	// 3
	// 9
	// 13
	// 7
	// 7
	// 6
	// 17
	// 20
	// 18
	// 15
}

func ExampleTree_Delete_noChildren() {
	t := makeTree()
	t.Delete(2)
	fmt.Println(t.Size())
	t.InOrderWalk(show)
	// Output:
	// 11
	// 3
	// 4
	// 6
	// 7
	// 7
	// 9
	// 13
	// 15
	// 17
	// 18
	// 20
}

func ExampleTree_Delete_root() {
	t := makeTree()
	t.Delete(15)
	fmt.Println(t.Size())
	t.InOrderWalk(show)
	// Output:
	// 11
	// 2
	// 3
	// 4
	// 6
	// 7
	// 7
	// 9
	// 13
	// 17
	// 18
	// 20
}

func ExampleTree_Delete_twoChildren() {
	t := makeTree()
	t.Delete(18)
	fmt.Println(t.Size())
	t.InOrderWalk(show)
	// Output:
	// 11
	// 2
	// 3
	// 4
	// 6
	// 7
	// 7
	// 9
	// 13
	// 15
	// 17
	// 20
}

func ExampleTree_Delete_leftChildOnly() {
	t := makeTree()
	t.Delete(13)
	fmt.Println(t.Size())
	t.InOrderWalk(show)
	// Output:
	// 11
	// 2
	// 3
	// 4
	// 6
	// 7
	// 7
	// 9
	// 15
	// 17
	// 18
	// 20
}

func ExampleTree_Delete_rightChildOnly() {
	t := makeTree()
	t.Insert(5)
	t.Delete(4)
	fmt.Println(t.Size())
	t.InOrderWalk(show)
	// Output:
	// 12
	// 2
	// 3
	// 5
	// 6
	// 7
	// 7
	// 9
	// 13
	// 15
	// 17
	// 18
	// 20
}

func ExampleTree_Delete_multipleItemsOfAllkinds() {
	t := makeTree()
	t.Insert(5)
	t.Delete(4)
	t.Delete(13)
	t.Delete(15)
	fmt.Println(t.Size())
	t.InOrderWalk(show)
	// Output:
	// 10
	// 2
	// 3
	// 5
	// 6
	// 7
	// 7
	// 9
	// 17
	// 18
	// 20
}
