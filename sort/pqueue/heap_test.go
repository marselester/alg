/*
A binary heap is a collection of keys arranged in a complete heap-ordered binary tree,
represented in level order in an array (not using the first entry):

0 1 2 3 4 5 6 7 8 9 10 11
- T S R P N O A E I  H  G
    ___T___
   S       R
  /  \    / \
 P    N   O A
/ \  / \
E I  H G
*/
package pqueue

import (
	"testing"
)

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestHeapSwim(t *testing.T) {
	tt := []struct {
		pq   []string
		i    int
		want []string
	}{
		{
			pq:   []string{"-", "A"},
			i:    1,
			want: []string{"-", "A"},
		},
		{
			pq:   []string{"-", "A", "B"},
			i:    2,
			want: []string{"-", "B", "A"},
		},
		{
			pq:   []string{"-", "S", "P", "R", "N", "T", "O", "A", "E", "I", "H", "G"},
			i:    5,
			want: []string{"-", "T", "S", "R", "N", "P", "O", "A", "E", "I", "H", "G"},
		},
	}
	for _, tc := range tt {
		h := Heap{pq: tc.pq}
		h.swim(tc.i)
		if !equal(h.pq, tc.want) {
			t.Errorf("swim(%d) got %v, want %v", tc.i, h.pq, tc.want)
		}
	}
}

func TestHeapSink(t *testing.T) {
	tt := []struct {
		pq   []string
		i    int
		want []string
	}{
		{
			pq:   []string{"-", "A"},
			i:    1,
			want: []string{"-", "A"},
		},
		{
			pq:   []string{"-", "A", "B"},
			i:    1,
			want: []string{"-", "B", "A"},
		},
		{
			pq:   []string{"-", "T", "H", "R", "P", "S", "O", "A", "E", "I", "N", "G"},
			i:    2,
			want: []string{"-", "T", "S", "R", "P", "N", "O", "A", "E", "I", "H", "G"},
		},
	}
	for _, tc := range tt {
		h := Heap{pq: tc.pq}
		h.sink(tc.i)
		if !equal(h.pq, tc.want) {
			t.Errorf("sink(%d) got %v, want %v", tc.i, h.pq, tc.want)
		}
	}
}

func TestHeapInsert(t *testing.T) {
	tt := []struct {
		pq   []string
		item string
		want []string
	}{
		{
			pq:   []string{"-"},
			item: "A",
			want: []string{"-", "A"},
		},
		{
			pq:   []string{"-", "A"},
			item: "B",
			want: []string{"-", "B", "A"},
		},
		{
			pq:   []string{"-", "T", "P", "R", "N", "H", "O", "A", "E", "I", "G"},
			item: "S",
			want: []string{"-", "T", "S", "R", "N", "P", "O", "A", "E", "I", "G", "H"},
		},
		{
			pq:   []string{"-", "T", "P", "R", "N", "H", "O", "A", "E", "I", "G"},
			item: "S",
			want: []string{"-", "T", "S", "R", "N", "P", "O", "A", "E", "I", "G", "H"},
		},
	}
	for _, tc := range tt {
		h := Heap{pq: tc.pq}
		h.Insert(tc.item)
		if !equal(h.pq, tc.want) {
			t.Errorf("Insert(%q) got %v, want %v", tc.item, h.pq, tc.want)
		}
	}
}

func TestHeapMax(t *testing.T) {
	tt := []struct {
		pq   []string
		max  string
		want []string
	}{
		{
			pq:   nil,
			max:  "",
			want: nil,
		},
		{
			pq:   []string{},
			max:  "",
			want: []string{},
		},
		{
			pq:   []string{"-"},
			max:  "",
			want: []string{"-"},
		},
		{
			pq:   []string{"-", "A"},
			max:  "A",
			want: []string{"-"},
		},
		{
			pq:   []string{"-", "B", "A"},
			max:  "B",
			want: []string{"-", "A"},
		},
		{
			pq:   []string{"-", "T", "S", "R", "N", "P", "O", "A", "E", "I", "G", "H"},
			max:  "T",
			want: []string{"-", "S", "P", "R", "N", "H", "O", "A", "E", "I", "G"},
		},
	}
	for _, tc := range tt {
		h := Heap{pq: tc.pq}
		if got := h.Max(); got != tc.max {
			t.Errorf("Max() = %q, want %q", got, tc.max)
		}
		if !equal(h.pq, tc.want) {
			t.Errorf("Max() got %v, want %v", h.pq, tc.want)
		}
	}
}

func TestHeap(t *testing.T) {
	h := NewHeap(1)
	h.Insert("P")
	h.Insert("Q")
	h.Insert("E")
	want := []string{"-", "Q", "P", "E"}
	if !equal(h.pq, want) {
		t.Errorf("Heap inserted P, Q, E got %v, want %v", h.pq, want)
	}

	h.Max()
	want = []string{"-", "P", "E"}
	if !equal(h.pq, want) {
		t.Errorf("Heap removed Q %v, want %v", h.pq, want)
	}

	h.Insert("X")
	h.Insert("A")
	h.Insert("M")
	want = []string{"-", "X", "M", "P", "A", "E"}
	if !equal(h.pq, want) {
		t.Errorf("Heap inserted X, A, M got %v, want %v", h.pq, want)
	}

	h.Max()
	want = []string{"-", "P", "M", "E", "A"}
	if !equal(h.pq, want) {
		t.Errorf("Heap removed X %v, want %v", h.pq, want)
	}

	h.Insert("P")
	h.Insert("L")
	h.Insert("E")
	want = []string{"-", "P", "P", "L", "A", "M", "E", "E"}
	if !equal(h.pq, want) {
		t.Errorf("Heap inserted P, L, E got %v, want %v", h.pq, want)
	}

	h.Max()
	want = []string{"-", "P", "M", "L", "A", "E", "E"}
	if !equal(h.pq, want) {
		t.Errorf("Heap removed P %v, want %v", h.pq, want)
	}
}
