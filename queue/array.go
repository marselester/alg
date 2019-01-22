package queue

// Array represents a queue of strings that is backed by array.
// Queue grows at a cost of allocating a new array and copying items there.
// When queue shrinks, it uses the same underlying array.
// The full array will be kept in memory until it is no longer referenced.
//
// Zero value is usable; initially queue has zero capacity.
// Note, operations are not concurrency safe.
type Array struct {
	items []string
}

// Enqueue adds an item to the end of the queue.
func (q *Array) Enqueue(item string) {
	q.items = append(q.items, item)
}

// Dequeue returns an item from the beginning of the queue.
func (q *Array) Dequeue() string {
	if len(q.items) == 0 {
		return ""
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v
}

// Size returns the number of items in the queue.
func (q *Array) Size() int {
	return len(q.items)
}
