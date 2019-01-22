package queue

// LinkedList represents a queue of strings based on linked-list data structure.
// The space required is always proportional to the size of of the collection.
// The time per operation is always independent of the size of the collection.
//
// The zero value for LinkedList is ready to use.
// Note, operations are not concurrency safe.
type LinkedList struct {
	// first is the node at the the beginning of the queue.
	first *node
	// last is the node at the end of the queue.
	last *node
	// n is the number of items in the queue.
	n int
}

type node struct {
	item string
	next *node
}

// Enqueue adds an item to the end of the queue.
func (q *LinkedList) Enqueue(item string) {
	newnode := node{item: item}
	if q.last != nil {
		q.last.next = &newnode
	}
	q.last = &newnode
	q.n++
	if q.first == nil {
		q.first = &newnode
	}
}

// Dequeue returns an item from the beginning of the queue.
func (q *LinkedList) Dequeue() string {
	if q.first == nil {
		return ""
	}

	v := q.first.item
	q.first = q.first.next
	q.n--
	if q.first == nil {
		q.last = nil
	}
	return v
}

// Size returns the number of items in the queue.
func (q *LinkedList) Size() int {
	return q.n
}
