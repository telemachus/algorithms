package algorithms_test

import (
	"testing"

	"github.com/telemachus/algorithms"
)

func BenchmarkInOrderIterativeWalk(b *testing.B) {
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
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.InOrderIterativeWalk(show)
	}
}

func BenchmarkInOrderRecursiveWalk(b *testing.B) {
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
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.InOrderWalk(show)
	}
}
