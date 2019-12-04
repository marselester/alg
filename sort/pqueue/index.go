package pqueue

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
	items []float64
}

// NewIndexMinHeap creates a binary heap of size n to prioritize min items.
func NewIndexMinHeap(n int) *IndexMinHeap {
	h := IndexMinHeap{
		pq:    make([]int, n+1),
		qp:    make([]int, n+1),
		items: make([]float64, n+1),
	}
	for i := 0; i <= n; i++ {
		h.qp[i] = -1
	}
	return &h
}

// Insert adds the new item and associates it with index i.
// Think of it as pq[i] = item.
func (h *IndexMinHeap) Insert(i int, item float64) {
	h.n++
	h.qp[i] = h.n
	h.pq[h.n] = i
	h.items[i] = item
	h.swim(h.n)
}

// Update changes the item associated with index i.
// Think of it as pq[i] = item.
func (h *IndexMinHeap) Update(i int, item float64) {
	h.items[i] = item
	h.swim(h.qp[i])
	h.sink(h.qp[i])
}

// Contains returns true if index i is associated with some item.
func (h *IndexMinHeap) Contains(i int) bool {
	return h.qp[i] != -1
}

// Min takes the smallest item off the top. Note, the first value is an index.
func (h *IndexMinHeap) Min() (int, float64) {
	if h.Size() == 0 {
		return -1, 0
	}

	indexOfMin := h.pq[1]
	min := h.items[indexOfMin]

	h.exchange(1, h.n)
	h.n--
	h.sink(1)

	h.items[h.pq[h.n+1]] = 0 // blank item
	h.qp[h.pq[h.n+1]] = -1

	return indexOfMin, min
}

// Size returns size of the heap.
func (h *IndexMinHeap) Size() int {
	return h.n
}

func (h *IndexMinHeap) greater(i, j int) bool {
	return h.items[h.pq[i]] > h.items[h.pq[j]]
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
