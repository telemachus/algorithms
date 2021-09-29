package algorithms_test

import (
	"testing"

	"git.sr.ht/~telemachus/algorithms"
)

func TestHeapLen(t *testing.T) {
	h := algorithms.NewHeap()
	if 0 != h.Len() {
		t.Errorf("expected %d; actual %d", 0, h.Len())
	}
}

func TestHeapRoot(t *testing.T) {
	h := algorithms.NewHeap()
	n, ok := h.Root()
	if 0 != n {
		t.Errorf("expected %d; actual %d", 0, n)
	}
	if false != ok {
		t.Errorf("expected %t, actual %t", false, ok)
	}
}
