package algorithms_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/telemachus/algorithms"
	"golang.org/x/tour/tree"
)

func BenchmarkRecursiveConcurrentWalk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		go algorithms.RecursiveConcurrentWalk(tree.New(2), ch)

		for n := range ch {
			fmt.Fprintln(io.Discard, n)
		}
	}
}

func BenchmarkIterativeConcurrentWalk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		go algorithms.IterativeConcurrentWalk(tree.New(2), ch)

		for n := range ch {
			fmt.Fprintln(io.Discard, n)
		}
	}
}

func BenchmarkMorrisIterativeConcurrentWalk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		go algorithms.MorrisIterativeConcurrentWalk(tree.New(2), ch)

		for n := range ch {
			fmt.Fprintln(io.Discard, n)
		}
	}
}
