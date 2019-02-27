package pqueue

/*
MaxHeap is a binary heap that can efficiently support priority-queue
operations: insert (log n), remove maximum (log n). The keys are stored in an array
such that each key is guaranteed to be >= the keys at two other positions.
This ordering represents a binary tree with edges from each key to the 2 smaller keys:
sequentially put the nodes in level order, with the root at index 1, its children at 2 and 3,
their children in positions 4, 5, 6, 7.

Parent of a[i] node is at a[i/2], children are at a[2*i] and a[2*i+1].
*/
type MaxHeap struct {
	// pq is a heap-ordered binary tree of string items.
	pq []string
}

// NewMaxHeap creates a binary heap of size n.
func NewMaxHeap(n int) *MaxHeap {
	h := MaxHeap{
		pq: make([]string, 0, n+1),
	}
	h.pq = append(h.pq, "-")
	return &h
}

// Insert adds the new item at the end of the array, and then swims up through the heap
// with that item to restore the heap condition.
func (h *MaxHeap) Insert(item string) {
	h.pq = append(h.pq, item)
	h.swim(len(h.pq) - 1)
}

// Max takes the largest item off the top, puts the item from the end of the heap at the top,
// decrements the size of the heap, and then sinks down through the heap with that item
// to restore the heap condition.
func (h *MaxHeap) Max() string {
	if len(h.pq) <= 1 {
		return ""
	}
	max := h.pq[1]
	h.pq[1] = h.pq[len(h.pq)-1]
	h.pq = h.pq[:len(h.pq)-1]
	h.sink(1)
	return max
}

// Size returns size of the heap.
func (h *MaxHeap) Size() int {
	return len(h.pq) - 1
}

// swim restores heap order by travelling from bottom up when
// a priority of some node i is increased (or a new node is added at the bottom of a heap).
// Exchange node with parent if it violates heap order (larger key than parent)
// until we reach a node with larger key, or the root.
func (h *MaxHeap) swim(i int) {
	var parent int
	for i > 1 {
		parent = i / 2
		if h.pq[i] > h.pq[parent] {
			h.pq[i], h.pq[parent] = h.pq[parent], h.pq[i]
		} else {
			break
		}
		i = parent
	}
}

// sink restores heap order by travelling down the heap when
// a priority of some node i is decreased. For example, a root node is replaced with a smaller key.
// Exchange node with the largest child if it violates heap order (smaller key than one or both of its children's keys)
// until we reach a node with both children smaller (or equal), or the bottom.
func (h *MaxHeap) sink(i int) {
	var child int
	for {
		// Find the largest child.
		child = 2 * i
		if child >= len(h.pq) {
			break
		}
		if child+1 < len(h.pq) && h.pq[child] < h.pq[child+1] {
			child++
		}

		if h.pq[i] < h.pq[child] {
			h.pq[i], h.pq[child] = h.pq[child], h.pq[i]
		} else {
			break
		}
		i = child
	}
}
