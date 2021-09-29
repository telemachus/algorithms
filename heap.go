package algorithms

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	h := &Heap{}
	h.data = make([]int, 0)
	return h
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
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

func (h *Heap) Insert(v int) {
	h.data = append(h.data, v)
	newNodeIndex := len(h.data) - 1
	for newNodeIndex > 0 && h.data[newNodeIndex] > h.data[parentIndex(newNodeIndex)] {
		h.data[parentIndex(newNodeIndex)], h.data[newNodeIndex] =
			h.data[newNodeIndex], h.data[parentIndex(newNodeIndex)]
		newNodeIndex = parentIndex(newNodeIndex)
	}
}

func (h *Heap) Delete() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	deletedValue := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data[len(h.data)-1] = 0
	h.data = h.data[:len(h.data)-1]
	trickleNodeIndex := 0

	for h.hasGreaterChild(trickleNodeIndex) {
		largerChildIndex := h.getLargerChildIndex(trickleNodeIndex)
		h.data[trickleNodeIndex], h.data[largerChildIndex] =
			h.data[largerChildIndex], h.data[trickleNodeIndex]
		trickleNodeIndex = largerChildIndex
	}

	return deletedValue, true
}

func (h *Heap) hasGreaterChild(n int) bool {
	lci := leftChildIndex(n)
	rci := rightChildIndex(n)

	return (h.validIndex(lci) && h.data[lci] > h.data[n]) || (h.validIndex(rci) && h.data[rci] > h.data[n])
}

func (h *Heap) getLargerChildIndex(n int) int {
	lci := leftChildIndex(n)
	rci := rightChildIndex(n)

	switch {
	case !h.validIndex(rci):
		return lci
	case h.data[rci] > h.data[lci]:
		return rci
	default:
		return lci
	}
}

func (h *Heap) validIndex(n int) bool {
	return n >= 0 && n < len(h.data)
}

func maxHeapify(data []int, i int) {
	leftIndex := leftChildIndex(i)
	rightIndex := rightChildIndex(i)
	var largestIndex int

	if leftIndex < len(data) && data[leftIndex] > data[i] {
		largestIndex = leftIndex
	} else {
		largestIndex = rightIndex
	}
	if rightIndex < len(data) && data[rightIndex] > data[largestIndex] {
		largestIndex = rightIndex
	}

	if largestIndex != index {
		data[i], data[largestIndex] = data[largestIndex], data[i]
		maxHeapify(data, largest)
	}
}
