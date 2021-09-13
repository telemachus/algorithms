package algorithms

// QuickselectL implements quickselect using Lomuto’s partition.
func QuickselectL(xs []int, low, high, nthLowest int) int {
	if len(xs) <= 0 || nthLowest > len(xs) {
		return -1
	}

	switch pivotIndex := lomutoPartition(xs, low, high); {
	case nthLowest > pivotIndex:
		return QuickselectL(xs, pivotIndex+1, high, nthLowest)
	case nthLowest < pivotIndex:
		return QuickselectL(xs, low, pivotIndex-1, nthLowest)
	default:
		return xs[nthLowest]
	}
}

// QuickselectH implements quickselect using Hoare’s partition.
func QuickselectH(xs []int, low, high, nthLowest int) int {
	if len(xs) <= 0 || nthLowest > len(xs) {
		return -1
	}

	switch pivotIndex := hoarePartition(xs, low, high); {
	case nthLowest > pivotIndex:
		return QuickselectH(xs, pivotIndex+1, high, nthLowest)
	case nthLowest < pivotIndex:
		return QuickselectH(xs, low, pivotIndex-1, nthLowest)
	default:
		return xs[nthLowest]
	}
}

// QuickselectYB implements quickselect using Stefan Nilsson’s partition.
func QuickselectYB(xs []int, nthLowest int) int {
	if len(xs) <= 0 || nthLowest > len(xs) {
		return -1
	}

	pivotValue := yourBasicPivot(xs)
	switch low, high := yourBasicPartition(xs, pivotValue); {
	case nthLowest <= low:
		QuicksortYB(xs[:low])
		return xs[nthLowest]
	default:
		QuicksortYB(xs[high:])
		return xs[nthLowest]
	}
}
