package queue

import "testing"

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestLinkedListEnqueue(t *testing.T) {
	q := LinkedList{}

	var items = []string{"a", "b", "c"}
	for _, v := range items {
		q.Enqueue(v)
	}

	var got []string
	for n := q.first; n != nil; n = n.next {
		got = append(got, n.item)
	}
	if !equal(got, items) {
		t.Errorf("Enqueue() got %q, want %q", got, items)
	}
}

func TestLinkedListEnqueueFirstItem(t *testing.T) {
	q := LinkedList{}

	q.Enqueue("fizz")
	want := node{item: "fizz"}
	if *q.first != want {
		t.Errorf("Enqueue(fizz) first is %v, want %v", *q.first, want)
	}
	if *q.last != want {
		t.Errorf("Enqueue(fizz) last is %v, want %v", *q.last, want)
	}
	if q.Size() != 1 {
		t.Errorf("Enqueue(fizz) count is %d, want 1", q.Size())
	}
}

func TestLinkedListDequeueEmpty(t *testing.T) {
	q := LinkedList{}

	got := q.Dequeue()
	want := ""
	if got != want {
		t.Errorf("Dequeue() = %q, want %q", got, want)
	}
}

func TestLinkedListDequeueLastItem(t *testing.T) {
	q := LinkedList{}

	q.Enqueue("fizz")
	got := q.Dequeue()
	want := "fizz"
	if got != want {
		t.Errorf("Dequeue() = %q, want %q", got, want)
	}
	if q.first != nil {
		t.Errorf("Dequeue() first must be nil")
	}
	if q.last != nil {
		t.Errorf("Dequeue() last must be nil")
	}
	if q.Size() != 0 {
		t.Errorf("Dequeue() count is %d, want 0", q.Size())
	}
}
