package algorithms_test

import (
	"testing"

	"git.sr.ht/~telemachus/algorithms"
)

func BenchmarkLomutoQuickselectShuffledSlice(b *testing.B) {
	xs := shuffledSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectL(xs, 0, len(xs)-1, 3567)
	}
}

func BenchmarkHoareQuickselectShuffledSlice(b *testing.B) {
	xs := shuffledSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectH(xs, 0, len(xs)-1, 3567)
	}
}

func BenchmarkYourBasicQuickselectShuffledSlice(b *testing.B) {
	xs := shuffledSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectYB(xs, 3567)
	}
}

func BenchmarkLomutoQuickselectEqualSlice(b *testing.B) {
	xs := equalSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectL(xs, 0, len(xs)-1, 3567)
	}
}

func BenchmarkHoareQuickselectEqualSlice(b *testing.B) {
	xs := equalSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectH(xs, 0, len(xs)-1, 3567)
	}
}

func BenchmarkYourBasicQuickselectEqualSlice(b *testing.B) {
	xs := equalSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectYB(xs, 3567)
	}
}

func BenchmarkLomutoQuickselectInOrderSlice(b *testing.B) {
	xs := inOrderSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectL(xs, 0, len(xs)-1, 3567)
	}
}

func BenchmarkHoareQuickselectInOrderSlice(b *testing.B) {
	xs := inOrderSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectH(xs, 0, len(xs)-1, 3567)
	}
}

func BenchmarkYourBasicQuickselectInOrderSlice(b *testing.B) {
	xs := inOrderSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectYB(xs, 3567)
	}
}

func BenchmarkLomutoQuickselectReversedSlice(b *testing.B) {
	xs := reversedSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectL(xs, 0, len(xs)-1, 3567)
	}
}

func BenchmarkHoareQuickselectReversedSlice(b *testing.B) {
	xs := reversedSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectH(xs, 0, len(xs)-1, 3567)
	}
}

func BenchmarkYourBasicQuickselectReversedSlice(b *testing.B) {
	xs := reversedSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithms.QuickselectYB(xs, 3567)
	}
}
