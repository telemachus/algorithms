package algorithms

type Heap []int

func NewHeap() Heap {
	h := make([]int, 0)
	return h
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Root() (int, bool) {
	if len(h) == 0 {
		return 0, false
	}
	return h[0], true
}

func leftChildIndex(n int) int {
	return (n * 2) + 1
}

func rightChildIndex(n int) int {
	return (n * 2) + 2
}

func parentIndex(n int) int {
	return (n - 1) / 2
}
