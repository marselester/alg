package stack

// LinkedList represents a stack of strings based on linked-list data structure.
// The space required is always proportional to the size of of the collection.
// The time per operation is always independent of the size of the collection.
//
// The zero value for LinkedList is ready to use.
// Note, operations are not concurrency safe.
type LinkedList struct {
	// first is the most recently added node (the top of the stack).
	first *node
	// n is the number of items in the stack.
	n int
}

type node struct {
	item string
	next *node
}

// Push adds an item to the top of the stack.
func (s *LinkedList) Push(item string) {
	first := node{item: item, next: s.first}
	s.first = &first
	s.n++
}

// Pop removes and returns the most recently added item.
// When stack is empty, empty string is returned.
func (s *LinkedList) Pop() string {
	if s.first == nil {
		return ""
	}
	v := s.first.item
	s.first = s.first.next
	s.n--
	return v
}

// Size returns the number of items in the stack.
func (s *LinkedList) Size() int {
	return s.n
}
