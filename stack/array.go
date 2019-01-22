package stack

// Array is a stack of strings that is backed by array.
// Stack grows at a cost of allocating a new array and copying items there.
// When stack shrinks, it uses the same underlying array.
// The full array will be kept in memory until it is no longer referenced.
//
// Zero value is usable; initially stack has zero capacity.
// Note, operations are not concurrency safe.
type Array struct {
	items []string
}

// NewArray creates a stack with given capacity.
// Negative capacity is ignored, zero value is used instead.
func NewArray(capacity int) *Array {
	if capacity < 0 {
		capacity = 0
	}
	return &Array{
		items: make([]string, 0, capacity),
	}
}

// Push adds an item to the top of the stack.
func (s *Array) Push(v string) {
	s.items = append(s.items, v)
}

// Pop removes and returns the most recently added item.
// When stack is empty, empty string is returned.
func (s *Array) Pop() string {
	i := len(s.items) - 1
	if i == -1 {
		return ""
	}
	v := s.items[i]
	s.items = s.items[:i]
	return v
}

// Size returns the number of items in the stack.
func (s *Array) Size() int {
	return len(s.items)
}
