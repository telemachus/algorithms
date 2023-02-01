package algorithms_test

import (
	"fmt"
	"testing"

	"github.com/telemachus/algorithms"
)

func TestNewBST(t *testing.T) {
	t.Parallel()

	tree := algorithms.NewBST()

	if tree == nil {
		t.Errorf("tree is nil, but it shouldn’t be")
	}

	if tree.Size() != 0 {
		t.Errorf("tree.Size() == %d; want 0", tree.Size())
	}

	want := "*algorithms.Tree"
	got := fmt.Sprintf("%T", tree)
	if got != want {
		t.Errorf("Sprintf(\"%%T\", tree) == %s; want %s", got, want)
	}
}

func TestTreeInsert(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(10)

	if tree.Size() != 4 {
		t.Errorf("tree.Size() == %d after inserting four items", tree.Size())
	}
}

func TestTreeSearch(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(10)

	got := tree.Search(-15)
	if got.Value() != -15 {
		t.Errorf("tree.Search(-15) == %d; want -15", got.Value())
	}

	got = tree.Search(7)
	if got != nil {
		t.Errorf("tree.Search(7) == %v; want nil", got)
	}

	got = tree.Search(10)
	if got.Count() != 2 {
		t.Errorf("got = tree.search(10); got.Count == %d; want 2", got.Count())
	}
}

func TestTreeMin(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(10)

	got := tree.Min()
	if got.Value() != -15 {
		t.Errorf("tree.Min() == %d; want -15", got.Value())
	}
}

func TestTreeMax(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(10)
	tree.Insert(1010)

	got := tree.Max()
	if got.Value() != 1010 {
		t.Errorf("tree.Max() == %d; want 1010", got.Value())
	}
}

func TestTreeSuccessor(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(15)
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(9)
	tree.Insert(18)
	tree.Insert(17)
	tree.Insert(20)

	node := tree.Search(7)
	got := tree.Successor(node)
	if got.Value() != 9 {
		t.Errorf("tree.Successor(%+v) == %d; want 9", got, got.Value())
	}

	node = tree.Search(20)
	got = tree.Successor(node)
	if got != nil {
		t.Errorf("tree.Successor(%+v) == %v; want nil", node, got)
	}
}

func TestTreePredecessor(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(15)
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(9)
	tree.Insert(18)
	tree.Insert(17)
	tree.Insert(20)

	node := tree.Search(15)
	got := tree.Predecessor(node)
	if got.Value() != 13 {
		t.Errorf("tree.Predecessor(%+v) == %d; want 13", got, got.Value())
	}

	node = tree.Search(4)
	got = tree.Predecessor(node)
	if got.Value() != 3 {
		t.Errorf("tree.Predecessor(%+v) == %d; want 3", node, got.Value())
	}

	node = tree.Search(9)
	got = tree.Predecessor(node)
	if got.Value() != 7 {
		t.Errorf("tree.Predecessor(%v) == %d; want 7", node, got.Value())
	}

	node = tree.Search(2)
	got = tree.Predecessor(node)
	if got != nil {
		t.Errorf("tree.Predecessor(%+v) == %v; want nil", node, got)
	}
}
