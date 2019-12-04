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

// IndexMinHeap is a binary heap that allows clients to refer to items on priority queue.
// The number of compares required is proportional to at most log n for insert, change priority,
// and remove the minimum.
type IndexMinHeap struct {
	// n is number of elements on priority queue.
	n int
	// pq is a binary heap using 1-based indexing.
	pq []int
	// qp is inverse: qp[pq[i]] = pq[qp[i]] = i.
	qp []int
	// items holds items with priorities.
	items []*Edge
}

// NewIndexMinHeap creates a binary heap of size n to prioritize min items.
func NewIndexMinHeap(n int) *IndexMinHeap {
	h := IndexMinHeap{
		pq:    make([]int, n+1),
		qp:    make([]int, n+1),
		items: make([]*Edge, n+1),
	}
	for i := 0; i <= n; i++ {
		h.qp[i] = -1
	}
	return &h
}

// Insert adds the new item and associates it with index i.
// Think of it as pq[i] = item.
func (h *IndexMinHeap) Insert(i int, item *Edge) {
	h.n++
	h.qp[i] = h.n
	h.pq[h.n] = i
	h.items[i] = item
	h.swim(h.n)
}

// Update changes the item associated with index i.
// Think of it as pq[i] = item.
func (h *IndexMinHeap) Update(i int, item *Edge) {
	h.items[i] = item
	h.swim(h.qp[i])
	h.sink(h.qp[i])
}

// Contains returns true if index i is associated with some item.
func (h *IndexMinHeap) Contains(i int) bool {
	return h.qp[i] != -1
}

// Min takes the smallest item off the top. Note, the first value is an index.
func (h *IndexMinHeap) Min() (int, *Edge) {
	if h.Size() == 0 {
		return -1, nil
	}

	indexOfMin := h.pq[1]
	min := h.items[indexOfMin]

	h.exchange(1, h.n)
	h.n--
	h.sink(1)

	h.items[h.pq[h.n+1]] = nil // blank item
	h.qp[h.pq[h.n+1]] = -1

	return indexOfMin, min
}

// Size returns size of the heap.
func (h *IndexMinHeap) Size() int {
	return h.n
}

func (h *IndexMinHeap) greater(i, j int) bool {
	return h.items[h.pq[i]].Weight > h.items[h.pq[j]].Weight
}

func (h *IndexMinHeap) exchange(i, j int) {
	swap := h.pq[i]
	h.pq[i] = h.pq[j]
	h.pq[j] = swap
	h.qp[h.pq[i]] = i
	h.qp[h.pq[j]] = j
}

func (h *IndexMinHeap) swim(k int) {
	for k > 1 && h.greater(k/2, k) {
		h.exchange(k, k/2)
		k = k / 2
	}
}

func (h *IndexMinHeap) sink(k int) {
	for 2*k <= h.n {
		j := 2 * k
		if j < h.n && h.greater(j, j+1) {
			j++
		}
		if !h.greater(k, j) {
			break
		}
		h.exchange(k, j)
		k = j
	}
}
