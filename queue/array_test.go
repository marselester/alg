package queue

import "testing"

func TestArrayEnqueue(t *testing.T) {
	q := Array{}

	var items = []string{"a", "b", "c"}
	for _, v := range items {
		q.Enqueue(v)
	}

	if !equal(q.items, items) {
		t.Errorf("Enqueue() got %q, want %q", q.items, items)
	}
}

func TestArrayDequeueEmpty(t *testing.T) {
	q := Array{}

	got := q.Dequeue()
	want := ""
	if got != want {
		t.Errorf("Dequeue() = %q, want %q", got, want)
	}
}

func TestArrayLayoutAfterEnqueueDequeue(t *testing.T) {
	q := Array{}

	q.Enqueue("fizz")
	if q.Size() != 1 {
		t.Errorf("Enqueue(fizz) count is %d, want 1", q.Size())
	}
	want := []string{"fizz"}
	if !equal(q.items, want) {
		t.Errorf("Enqueue(fizz) got items %v, want %v", q.items, want)
	}

	got := q.Dequeue()
	if got != "fizz" {
		t.Errorf("Dequeue() got %v, want fizz", got)
	}
	if q.Size() != 0 {
		t.Errorf("Dequeue() count is %d, want 0", q.Size())
	}
	if !equal(q.items, nil) {
		t.Errorf("Dequeue() got items %v, want none", q.items)
	}
}
