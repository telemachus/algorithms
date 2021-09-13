package algorithms

// InsertionSort sorts a slice in place.
func InsertionSort(xs []int) {
	// Loop invariant 1: xs[:j-1] contains the items originally in that
	// subslice but in sorted order.
	for j := 1; j < len(xs); j++ {
		key := xs[j]
		i := j - 1

		// Loop invariant 2: xs[i:j] accepts items only <= key.
		for i >= 0 && xs[i] > key {
			xs[i+1] = xs[i]
			i--
		}
		xs[i+1] = key
	}
}
