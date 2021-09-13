package algorithms_test

import (
	"math/rand"
	"sort"
	"testing"

	"git.sr.ht/~telemachus/algorithms"
)

func shuffledSlice() sort.IntSlice {
	var xs sort.IntSlice
	xs = make([]int, 10000)
	rand.Shuffle(len(xs), func(i, j int) {
		xs[i], xs[j] = xs[j], xs[i]
	})

	return xs
}

func equalSlice() sort.IntSlice {
	var xs sort.IntSlice
	xs = make([]int, 10000)
	for i := 0; i < 10000; i++ {
		xs[i] = 1
	}

	return xs
}

func inOrderSlice() sort.IntSlice {
	var xs sort.IntSlice
	xs = make([]int, 10000)
	for i := 0; i < 10000; i++ {
		xs[i] = i
	}

	return xs
}

func reversedSlice() sort.IntSlice {
	var xs sort.IntSlice
	xs = make([]int, 10000)
	for i, j := 0, 10000; i < 10000; i, j = i+1, j-1 {
		xs[i] = j
	}

	return xs
}

func BenchmarkLomutoQuicksortShuffledSlice(b *testing.B) {
	xs := shuffledSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortL(xs, 0, len(xs)-1)
	}
}

func BenchmarkHoareQuicksortShuffledSlice(b *testing.B) {
	xs := shuffledSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortH(xs, 0, len(xs)-1)
	}
}

func BenchmarkStdlibSortShuffledSlice(b *testing.B) {
	xs := shuffledSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		xs.Sort()
	}
}

func BenchmarkYourBasicQuicksortShuffledSlice(b *testing.B) {
	xs := shuffledSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortYB(xs)
	}
}

func BenchmarkLomutoQuicksortEqualSlice(b *testing.B) {
	xs := equalSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortL(xs, 0, len(xs)-1)
	}
}

func BenchmarkHoareQuicksortEqualSlice(b *testing.B) {
	xs := equalSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortH(xs, 0, len(xs)-1)
	}
}

func BenchmarkStdlibSortEqualSlice(b *testing.B) {
	xs := equalSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		xs.Sort()
	}
}

func BenchmarkYourBasicQuicksortEqualSlice(b *testing.B) {
	xs := equalSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortYB(xs)
	}
}

func BenchmarkLomutoQuicksortInOrderSlice(b *testing.B) {
	xs := inOrderSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortL(xs, 0, len(xs)-1)
	}
}

func BenchmarkHoareQuicksortInOrderSlice(b *testing.B) {
	xs := inOrderSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortH(xs, 0, len(xs)-1)
	}
}

func BenchmarkStdlibSortInOrderSlice(b *testing.B) {
	xs := inOrderSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		xs.Sort()
	}
}

func BenchmarkYourBasicQuicksortInOrderSlice(b *testing.B) {
	xs := inOrderSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortYB(xs)
	}
}

func BenchmarkLomutoQuicksortReversedSlice(b *testing.B) {
	xs := reversedSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortL(xs, 0, len(xs)-1)
	}
}

func BenchmarkHoareQuicksortReversedSlice(b *testing.B) {
	xs := reversedSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortH(xs, 0, len(xs)-1)
	}
}

func BenchmarkStdlibSortReversedSlice(b *testing.B) {
	xs := reversedSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		xs.Sort()
	}
}

func BenchmarkYourBasicQuicksortReversedSlice(b *testing.B) {
	xs := reversedSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuicksortYB(xs)
	}
}
