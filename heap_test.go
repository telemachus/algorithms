package algorithms_test

import (
	"testing"

	"git.sr.ht/~telemachus/algorithms"
)

func TestHeapLen(t *testing.T) {
	h := algorithms.NewHeap()
	if 0 != h.Len() {
		t.Errorf("expected 0; actual %d", h.Len())
	}

	h.Insert(10)
	if 1 != h.Len() {
		t.Errorf("expected 1; actual %d", h.Len())
	}
}

func TestInsert(t *testing.T) {
	h := algorithms.NewHeap()

	h.Insert(1)
	h.Insert(20)
	h.Insert(3)
	h.Insert(-10)
	h.Insert(400)
	if 5 != h.Len() {
		t.Errorf("expected 5; actual %d", h.Len())
	}
	actual, ok := h.Peek()
	if 400 != actual {
		t.Errorf("expected 400; actual %d", actual)
	}
	if true != ok {
		t.Errorf("expected 400; actual %t", ok)
	}
}

func TestDelete(t *testing.T) {
	h := algorithms.NewHeap()

	h.Insert(20)
	h.Insert(-10)
	h.Insert(400)

	first, _ := h.Delete()
	if 400 != first {
		t.Errorf("expected 400; actual %d", first)
	}

	second, _ := h.Delete()
	if 20 != second {
		t.Errorf("expected 20; actual %d", second)
	}

	third, ok := h.Delete()
	if -10 != third {
		t.Errorf("expected -10; actual %d", third)
	}
	if true != ok {
		t.Errorf("expected true; actual %t", ok)
	}

	_, ok = h.Delete()
	if false != ok {
		t.Errorf("expected false; actual %t", ok)
	}
}

func TestHeapPeek(t *testing.T) {
	h := algorithms.NewHeap()
	n, ok := h.Peek()
	if 0 != n {
		t.Errorf("expected 0; actual %d", n)
	}
	if false != ok {
		t.Errorf("expected false, actual %t", ok)
	}

	h.Insert(20)
	h.Insert(-10)
	h.Insert(400)

	top, ok := h.Peek()
	if 400 != top {
		t.Errorf("expected 400; actual %d", top)
	}
	if true != ok {
		t.Errorf("expected true, actual %t", ok)
	}
	if 3 != h.Len() {
		t.Errorf("expected 3; actual %d", h.Len())
	}
}
