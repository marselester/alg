package stack

import "testing"

func TestLinkedListPopEmpty(t *testing.T) {
	s := LinkedList{}

	got := s.Pop()
	want := ""
	if got != want {
		t.Errorf("Pop() = %q, want %q", got, want)
	}
}

func TestLinkedListLayoutAfterPushPush(t *testing.T) {
	s := LinkedList{}

	s.Push("fizz")
	if s.Size() != 1 {
		t.Errorf("Push(fizz) count is %d, want 1", s.Size())
	}
	if s.first.item != "fizz" {
		t.Errorf("Push(fizz) got item %v, want fizz", s.first.item)
	}

	s.Push("bazz")
	if s.Size() != 2 {
		t.Errorf("Push(bazz) count is %d, want 2", s.Size())
	}
	want := []string{"bazz", "fizz"}
	got := []string{s.first.item, s.first.next.item}
	if !equal(got, want) {
		t.Errorf("Push(bazz) got items %v, want %v", got, want)
	}
}

func TestLinkedListLayoutAfterPushPop(t *testing.T) {
	s := LinkedList{}

	s.Push("fizz")
	got := s.Pop()
	if got != "fizz" {
		t.Errorf("Pop() got %v, want fizz", got)
	}
	if s.Size() != 0 {
		t.Errorf("Pop() count is %d, want 0", s.Size())
	}
	if s.first != nil {
		t.Errorf("Pop() has node %v, want nil", s.first)
	}
}
