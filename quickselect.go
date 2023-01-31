package algorithms

// QuickselectL implements quickselect using Lomuto’s partition.
func QuickselectL(xs []int, low, high, nthLowest int) int {
	if len(xs) < 1 || nthLowest > len(xs) {
		return -1
	}

	QuicksortL(xs, low, high)

	return xs[nthLowest]
}

// QuickselectH implements quickselect using Hoare’s partition.
func QuickselectH(xs []int, low, high, nthLowest int) int {
	if len(xs) < 1 || nthLowest > len(xs) {
		return -1
	}

	QuicksortH(xs, low, high)

	return xs[nthLowest]
}

// QuickselectYB implements quickselect using Stefan Nilsson’s partition.
func QuickselectYB(xs []int, nthLowest int) int {
	if len(xs) < 1 || nthLowest > len(xs) {
		return -1
	}

	QuicksortYB(xs)

	return xs[nthLowest]
}
