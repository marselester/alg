package mst

// MinHeap is a binary heap that can efficiently support priority-queue
// operations: insert (log n), remove minimum (log n).
type MinHeap struct {
	pq []*Edge
}

// NewMinHeap creates a binary heap of size n to prioritize min items.
func NewMinHeap(n int) *MinHeap {
	h := MinHeap{
		pq: make([]*Edge, 0, n+1),
	}
	h.pq = append(h.pq, nil)
	return &h
}

// Insert adds the new item at the end of the array, and then swims up through the heap
// with that item to restore the heap condition.
func (h *MinHeap) Insert(item *Edge) {
	h.pq = append(h.pq, item)
	h.swim(len(h.pq) - 1)
}

// Min takes the lightest item off the top, puts the item from the end of the heap at the top,
// decrements the size of the heap, and then sinks down through the heap with that item
// to restore the heap condition.
func (h *MinHeap) Min() *Edge {
	if len(h.pq) <= 1 {
		return nil
	}
	min := h.pq[1]
	h.pq[1] = h.pq[len(h.pq)-1]
	h.pq = h.pq[:len(h.pq)-1]
	h.sink(1)
	return min
}

// Size returns size of the heap.
func (h *MinHeap) Size() int {
	return len(h.pq) - 1
}

func (h *MinHeap) swim(i int) {
	var parent int
	for i > 1 {
		parent = i / 2
		if h.pq[i].Weight < h.pq[parent].Weight {
			h.pq[i], h.pq[parent] = h.pq[parent], h.pq[i]
		} else {
			break
		}
		i = parent
	}
}

func (h *MinHeap) sink(i int) {
	var child int
	for {
		// Find the smallest child.
		child = 2 * i
		if child >= len(h.pq) {
			break
		}
		if child+1 < len(h.pq) && h.pq[child].Weight > h.pq[child+1].Weight {
			child++
		}

		if h.pq[i].Weight > h.pq[child].Weight {
			h.pq[i], h.pq[child] = h.pq[child], h.pq[i]
		} else {
			break
		}
		i = child
	}
}
