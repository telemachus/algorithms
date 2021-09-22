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
