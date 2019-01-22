package stack

import (
	"testing"
)

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

func TestArrayPopEmpty(t *testing.T) {
	s := Array{}

	got := s.Pop()
	want := ""
	if got != want {
		t.Errorf("Pop() = %q, want %q", got, want)
	}
}

func TestArrayLayoutAfterPushPush(t *testing.T) {
	s := Array{}

	s.Push("fizz")
	if s.Size() != 1 {
		t.Errorf("Push(fizz) count is %d, want 1", s.Size())
	}
	want := []string{"fizz"}
	if !equal(s.items, want) {
		t.Errorf("Push(fizz) got items %v, want %v", s.items, want)
	}

	s.Push("bazz")
	if s.Size() != 2 {
		t.Errorf("Push(bazz) count is %d, want 2", s.Size())
	}
	want = []string{"fizz", "bazz"}
	if !equal(s.items, want) {
		t.Errorf("Push(bazz) got items %v, want %v", s.items, want)
	}
}

func TestArrayLayoutAfterPushPop(t *testing.T) {
	s := Array{}

	s.Push("fizz")
	if s.Size() != 1 {
		t.Errorf("Push(fizz) count is %d, want 1", s.Size())
	}
	want := []string{"fizz"}
	if !equal(s.items, want) {
		t.Errorf("Push(fizz) got items %v, want %v", s.items, want)
	}

	got := s.Pop()
	if got != "fizz" {
		t.Errorf("Pop() got %v, want fizz", got)
	}
	if s.Size() != 0 {
		t.Errorf("Pop() count is %d, want 0", s.Size())
	}
	if !equal(s.items, nil) {
		t.Errorf("Pop() got items %v, want none", s.items)
	}
}

func TestNewArrayNegativeCap(t *testing.T) {
	s := NewArray(-1)

	wantCap := 0
	if cap(s.items) != wantCap {
		t.Errorf("NewArray(-1) cap = %d, want %d", cap(s.items), wantCap)
	}
}

func TestNewArray(t *testing.T) {
	s := NewArray(10)

	s.Push("fizz")
	got := s.Pop()
	want := "fizz"
	if got != want {
		t.Errorf("Pop() = %q, want %q", got, want)
	}

	wantCap := 10
	if cap(s.items) != wantCap {
		t.Errorf("NewArray(10) cap = %d, want %d", cap(s.items), wantCap)
	}
}

func TestNewArrayZero(t *testing.T) {
	s := NewArray(0)

	s.Push("fizz")
	got := s.Pop()
	want := "fizz"
	if got != want {
		t.Errorf("Pop() = %q, want %q", got, want)
	}
}
