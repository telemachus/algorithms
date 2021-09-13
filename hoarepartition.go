package algorithms

func hoarePartition(xs []int, low, high int) int {
	pivotValue := xs[(high+low)/2]
	i := low - 1
	j := high + 1

	for {
		for i = i + 1; xs[i] < pivotValue; i++ {
		}
		for j = j - 1; xs[j] > pivotValue; j-- {
		}

		// The loops below are a more literal version of Hoare’s partition,
		// but they are far less idiomatic Go.
		//
		// for {
		// 	i++
		// 	if xs[i] >= pivotValue {
		// 		break
		// 	}
		// }
		// for {
		// 	j--
		// 	if xs[j] <= pivotValue {
		// 		break
		// 	}
		// }

		if i < j {
			xs[i], xs[j] = xs[j], xs[i]
		} else {
			return j
		}
	}
}
