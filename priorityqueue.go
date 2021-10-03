package algorithms

import "math"

type Item struct {
	Name     string
	Priority int
}

type PriorityQueue struct {
	data []*Item
}

func NewPQ() *PriorityQueue {
	return &PriorityQueue{}
}

func (pq *PriorityQueue) Init(items []*Item) {
	pq.data = make([]*Item, len(items))
	copy(pq.data, items)
	for i := (len(pq.data) - 1) / 2; i >= 0; i-- {
		pq.maxHeapify(i)
	}
}

func (pq *PriorityQueue) Max() (*Item, bool) {
	if len(pq.data) == 0 {
		return nil, false
	}
	return pq.data[0], true
}

func (pq *PriorityQueue) RemoveMax() (*Item, bool) {
	dataCount := len(pq.data)
	if dataCount == 0 {
		return nil, false
	}
	dataCount--
	max := pq.data[0]
	pq.data[0] = pq.data[dataCount]
	pq.data[dataCount] = nil
	pq.data = pq.data[:dataCount]
	pq.maxHeapify(0)
	return max, true
}

func (pq *PriorityQueue) isValid(i int) bool {
	return i >= 0 && i < len(pq.data)
}

func (pq *PriorityQueue) findMaxIndex(i int) int {
	left := leftChildIndex(i)
	right := rightChildIndex(i)
	var largest int

	if pq.isValid(left) && pq.data[left].Priority > pq.data[i].Priority {
		largest = left
	} else {
		largest = i
	}
	if pq.isValid(right) && pq.data[right].Priority > pq.data[largest].Priority {
		largest = right
	}
	return largest
}

func (pq *PriorityQueue) maxHeapify(i int) {
	largest := pq.findMaxIndex(i)
	for largest != i {
		pq.data[i], pq.data[largest] = pq.data[largest], pq.data[i]
		i = largest
		largest = pq.findMaxIndex(i)
	}
}

func (pq *PriorityQueue) increaseKey(i, key int) {
	if key <= pq.data[i].Priority {
		return
	}
	pq.data[i].Priority = key
	for i > 0 && pq.data[parentIndex(i)].Priority < pq.data[i].Priority {
		pq.data[i], pq.data[parentIndex(i)] = pq.data[parentIndex(i)], pq.data[i]
		i = parentIndex(i)
	}
}

func (pq *PriorityQueue) Insert(n string, p int) {
	item := &Item{Name: n, Priority: math.MinInt}
	pq.data = append(pq.data, item)
	pq.increaseKey(len(pq.data)-1, p)
}
