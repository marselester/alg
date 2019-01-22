package queue

// RingBuffer (circular queue) is a FIFO data structure of a fixed size
// https://en.wikipedia.org/wiki/Circular_buffer.
// When the buffer is empty, the consumer waits until data is deposited;
// when the buffer is full, the producer waits to deposit data.
// It is useful for transferring data between asynchronous processes or
// for storing log files.
//
// This implementation uses an array representation.
// Zero value is unusable (it has zero capacity), please use NewRingBuffer.
// Note, operations are not concurrency safe.
type RingBuffer struct {
	items []string
	// n indicates number of items in the queue.
	n int
	// writePos is an index in the array where a new item will be written at (enqueue).
	writePos int
	// readPos is an index in the array where an item will be read from (dequeue).
	readPos int
}

// NewRingBuffer returns a ring buffer of a fixed size capacity.
func NewRingBuffer(capacity int) *RingBuffer {
	if capacity < 0 {
		capacity = 0
	}
	return &RingBuffer{
		items: make([]string, capacity),
	}
}

// Enqueue adds an item to the end of the queue.
// When queue is full, it returns false.
func (r *RingBuffer) Enqueue(item string) bool {
	// Queue is full.
	if r.n == len(r.items) {
		return false
	}
	r.items[r.writePos] = item
	r.writePos = next(r.writePos, len(r.items))
	r.n++
	return true
}

// Dequeue returns an item from the beginning of the queue.
func (r *RingBuffer) Dequeue() (string, bool) {
	// Queue is empty.
	if r.n == 0 {
		return "", false
	}
	v := r.items[r.readPos]
	r.items[r.readPos] = "" // Clean up so it's easier to test.
	r.readPos = next(r.readPos, len(r.items))
	r.n--
	return v, true
}

// Size returns the number of items in the queue.
func (r *RingBuffer) Size() int {
	return r.n
}

// next moves cursor i in a circle with fixed capacity.
// For example, a circle has capacity of 3 elements, then
// next positions of a cursor are 0 -> 1 -> 2 -> 0 -> 1 and so on.
// When invariant is violated (index >= capacity), -1 is returned.
func next(i, capacity int) int {
	if i >= capacity {
		return -1
	}

	i++
	if i == capacity {
		return 0
	}
	return i
}
