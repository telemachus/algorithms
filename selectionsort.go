package algorithms

func SelectionSort(xs []int) {
	for i := 0; i < len(xs)-1; i++ {
		minIndex := findMinIndexFrom(xs, i)
		xs[i], xs[minIndex] = xs[minIndex], xs[i]
	}
}

func findMinIndexFrom(xs []int, n int) int {
	minIndex := n
	for i := n + 1; i < len(xs); i++ {
		if xs[i] < xs[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}
