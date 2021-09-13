package algorithms

// QuicksortL implements quicksort using Lomuto’s partition.
func QuicksortL(xs []int, low, high int) {
	if high-low <= 0 {
		return
	}

	pivotIndex := lomutoPartition(xs, low, high)
	QuicksortL(xs, low, pivotIndex-1)
	QuicksortL(xs, pivotIndex+1, high)
}

// QuicksortH implements quicksort using Hoare’s partition.
func QuicksortH(xs []int, low, high int) {
	if high-low <= 0 {
		return
	}

	pivotIndex := hoarePartition(xs, low, high)
	QuicksortH(xs, low, pivotIndex)
	QuicksortH(xs, pivotIndex+1, high)
}

// QuicksortYB implements quicksort using Stefan Nilsson’s partition.
// See the following links:
// - https://yourbasic.org/golang/quicksort-optimizations
// - https://yourbasic.org/about
func QuicksortYB(xs []int) {
	if len(xs) < 20 {
		InsertionSort(xs)
		return
	}

	pivotIndex := yourBasicPivot(xs)
	low, high := yourBasicPartition(xs, pivotIndex)
	QuicksortYB(xs[:low])
	QuicksortYB(xs[high:])
}
