package queue

import "testing"

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(3)

	var want = []string{"", "", ""}
	if !equal(rb.items, want) {
		t.Errorf("NewRingBuffer() items %q, want %q", rb.items, want)
	}

	rb.Enqueue("1")
	want = []string{"1", "", ""}
	if !equal(rb.items, want) {
		t.Errorf("Enqueue(1) items %q, want %q", rb.items, want)
	}

	rb.Enqueue("2")
	want = []string{"1", "2", ""}
	if !equal(rb.items, want) {
		t.Errorf("Enqueue(1) items %q, want %q", rb.items, want)
	}

	if got, ok := rb.Dequeue(); !ok || got != "1" {
		t.Errorf("Dequeue() = %q, %v; want 1, true", got, ok)
	}
	want = []string{"", "2", ""}
	if !equal(rb.items, want) {
		t.Errorf("Dequeue() items %q, want %q", rb.items, want)
	}

	rb.Enqueue("3")
	want = []string{"", "2", "3"}
	if !equal(rb.items, want) {
		t.Errorf("Enqueue(3) items %q, want %q", rb.items, want)
	}

	if got, ok := rb.Dequeue(); !ok || got != "2" {
		t.Errorf("Dequeue() = %q, %v; want 2, true", got, ok)
	}
	want = []string{"", "", "3"}
	if !equal(rb.items, want) {
		t.Errorf("Dequeue() items %q, want %q", rb.items, want)
	}

	rb.Enqueue("4")
	want = []string{"4", "", "3"}
	if !equal(rb.items, want) {
		t.Errorf("Enqueue(4) items %q, want %q", rb.items, want)
	}
}

func TestNewRingBufferZero(t *testing.T) {
	rb := NewRingBuffer(0)

	if ok := rb.Enqueue("fizz"); ok {
		t.Errorf("Enqueue(fizz) = %v, want false", ok)
	}
}

func TestNewRingBufferNegativeCap(t *testing.T) {
	rb := NewRingBuffer(-1)

	wantCap := 0
	if cap(rb.items) != wantCap {
		t.Errorf("NewRingBuffer(-1) cap = %d, want %d", cap(rb.items), wantCap)
	}
}

func TestNext(t *testing.T) {
	tt := []struct {
		i        int
		capacity int
		want     int
	}{
		{0, 3, 1},
		{1, 3, 2},
		{2, 3, 0},
		{2, 4, 3},
		{0, 0, -1},
		{1, 0, -1},
	}
	for _, tc := range tt {
		got := next(tc.i, tc.capacity)
		if got != tc.want {
			t.Errorf("next(%d, %d) = %d, want %d", tc.i, tc.capacity, got, tc.want)
		}
	}
}
