package algorithms

func BinarySearch(xs []int, wanted int) int {
	min := 0
	max := len(xs) - 1
	var guess int

	for min <= max {
		guess = (min + max) / 2
		switch {
		case xs[guess] < wanted:
			min = guess + 1
		case xs[guess] > wanted:
			max = guess - 1
		default:
			return guess
		}
	}
	return -1
}
