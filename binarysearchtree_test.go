package algorithms_test

import (
	"fmt"
	"testing"

	"git.sr.ht/~telemachus/algorithms"
)

func TestNewBST(t *testing.T) {
	tree := algorithms.NewBST()

	if tree == nil {
		t.Errorf("tree is nil, but it shouldn’t be")
	}
	if 0 != tree.Size() {
		t.Errorf("expected tree.Size() to be 0; actual: %d", tree.Size())
	}
	expected := "*algorithms.Tree"
	actual := fmt.Sprintf("%T", tree)
	if expected != actual {
		t.Errorf("expected %s; actual %s", expected, actual)
	}
}

func TestTreeInsert(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(10)

	if 4 != tree.Size() {
		t.Errorf("inserted three items; actual tree.Size() is %d", tree.Size())
	}
}

func TestTreeSearch(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(10)

	actual := tree.Search(-15)
	if -15 != actual.Value() {
		t.Errorf("expected %d; actual %d", -15, actual.Value())
	}

	actual = tree.Search(7)
	if nil != actual {
		t.Errorf("expected %#v; actual %#v", nil, actual)
	}

	actual = tree.Search(10)
	if 2 != actual.Count() {
		t.Errorf("expected result to have 10 twice; actual %d", actual.Count())
	}
}

func TestTreeMin(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(10)

	actual := tree.Min()
	if -15 != actual.Value() {
		t.Errorf("expected a minimum of %d; actual %d", -15, actual.Value())
	}
}

func TestTreeMax(t *testing.T) {
	tree := algorithms.NewBST()
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(10)
	tree.Insert(1010)

	actual := tree.Max()
	if 1010 != actual.Value() {
		t.Errorf("expected a maximum of %d; actual %d", 1010, actual.Value())
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
	actual := tree.Successor(node)
	if 9 != actual.Value() {
		t.Errorf("expected 9 as successor; actual %d", actual.Value())
	}

	node = tree.Search(20)
	actual = tree.Successor(node)
	if nil != actual {
		t.Errorf("expected nil as successor; actual %#v", actual)
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
	actual := tree.Predecessor(node)
	if 13 != actual.Value() {
		t.Errorf("expected 15 as predecessor; actual %d", actual.Value())
	}

	node = tree.Search(4)
	actual = tree.Predecessor(node)
	if 3 != actual.Value() {
		t.Errorf("expected 3 as predecessor; actual %d", actual.Value())
	}

	node = tree.Search(9)
	actual = tree.Predecessor(node)
	if 7 != actual.Value() {
		t.Errorf("expected 7 as predecessor; actual %d", actual.Value())
	}

	node = tree.Search(2)
	actual = tree.Predecessor(node)
	if nil != actual {
		t.Errorf("expected nil as predecessor; actual %#v", actual)
	}
}
