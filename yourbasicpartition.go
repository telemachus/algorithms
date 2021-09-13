package algorithms

import (
	"math/rand"
)

func yourBasicPivot(xs []int) int {
	n := len(xs)
	return yourBasicMedian(xs[rand.Intn(n)],
		xs[rand.Intn(n)],
		xs[rand.Intn(n)])
}

func yourBasicMedian(a, b, c int) int {
	if a < b {
		switch {
		case b < c:
			return b
		case a < c:
			return c
		default:
			return a
		}
	}
	switch {
	case a < c:
		return a
	case b < c:
		return c
	default:
		return b
	}
}

// Partition reorders the elements of xs so that:
// - all elements in xs[:low] are less than pivotValue,
// - all elements in xs[low:high] are equal to pivotValue,
// - all elements in xs[high:] are greater than pivotValue.
func yourBasicPartition(xs []int, pivotValue int) (low, high int) {
	low, high = 0, len(xs)
	for mid := 0; mid < high; {
		// Invariant:
		//  - xs[:low] < pivotValue
		//  - xs[low:mid] = pivotValue
		//  - xs[mid:high] are unknown
		//  - xs[high:] > pivotValue
		//
		//     < pivotValue  = pivotValue  unknown       > pivotValue
		//     ----------------------------------------------------
		// xs: |            |            |a            |           |
		//     ----------------------------------------------------
		//                  ^            ^             ^
		//                  low          mid           high
		switch x := xs[mid]; {
		case x < pivotValue:
			xs[mid] = xs[low]
			xs[low] = x
			low++
			mid++
		case x == pivotValue:
			mid++
		default: // x > pivotValue
			xs[mid] = xs[high-1]
			xs[high-1] = x
			high--
		}
	}
	return low, high
}
