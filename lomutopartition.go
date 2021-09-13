package algorithms

func lomutoPartition(xs []int, low, high int) int {
	pivotValue := xs[high]
	i := low - 1

	for j := low; j < high; j++ {
		if xs[j] <= pivotValue {
			i++
			xs[i], xs[j] = xs[j], xs[i]
		}
	}

	i++
	xs[i], xs[high] = xs[high], xs[i]
	return i
}
