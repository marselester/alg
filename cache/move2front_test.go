package cache

import (
	"fmt"
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

// cachedKeyVal returns a list of key-value pairs from linked list for testing.
func cachedKeyVal(first *node) []string {
	var kv []string
	for n := first; n != nil; n = n.next {
		kv = append(kv, fmt.Sprintf("%s:%s", n.key, n.value))
	}
	return kv
}

func TestMoveToFrontSet(t *testing.T) {
	var tt = []struct {
		cache MoveToFront
		key   string
		value []byte
		want  []string
	}{
		{
			cache: MoveToFront{},
			key:   "a",
			value: []byte("A"),
			want:  []string{"a:A"},
		},
		{
			cache: MoveToFront{
				first: &node{
					key:   "b",
					value: []byte("B"),
				},
			},
			key:   "a",
			value: []byte("A"),
			want:  []string{"a:A", "b:B"},
		},
		{
			cache: MoveToFront{
				first: &node{
					key:   "a",
					value: []byte("A"),
				},
			},
			key:   "a",
			value: []byte("AAA"),
			want:  []string{"a:AAA"},
		},
		{
			cache: MoveToFront{
				first: &node{
					key:   "b",
					value: []byte("B"),
					next: &node{
						key:   "a",
						value: []byte("A"),
					},
				},
			},
			key:   "a",
			value: []byte("AAA"),
			want:  []string{"a:AAA", "b:B"},
		},
		{
			cache: MoveToFront{
				first: &node{
					key:   "a",
					value: []byte("A"),
					next: &node{
						key:   "b",
						value: []byte("B"),
					},
				},
			},
			key:   "a",
			value: []byte("AAA"),
			want:  []string{"a:AAA", "b:B"},
		},
	}

	for _, tc := range tt {
		tc.cache.Set(tc.key, tc.value)
		got := cachedKeyVal(tc.cache.first)
		if !equal(got, tc.want) {
			t.Errorf("Set(%q, %q) = %v, want %v", tc.key, tc.value, got, tc.want)
		}
	}
}

func TestMoveToFrontGet(t *testing.T) {
	var tt = []struct {
		cache MoveToFront
		key   string
		want  string
	}{
		{
			cache: MoveToFront{},
			key:   "a",
			want:  "",
		},
		{
			cache: MoveToFront{
				first: &node{
					key:   "b",
					value: []byte("B"),
				},
			},
			key:  "a",
			want: "",
		},
		{
			cache: MoveToFront{
				first: &node{
					key:   "b",
					value: []byte("B"),
				},
			},
			key:  "b",
			want: "B",
		},
		{
			cache: MoveToFront{
				first: &node{
					key:   "b",
					value: []byte("B"),
					next: &node{
						key:   "a",
						value: []byte("A"),
					},
				},
			},
			key:  "a",
			want: "A",
		},
		{
			cache: MoveToFront{
				first: &node{
					key:   "a",
					value: []byte("A"),
					next: &node{
						key:   "b",
						value: []byte("B"),
					},
				},
			},
			key:  "a",
			want: "A",
		},
	}

	for _, tc := range tt {
		got := string(tc.cache.Get(tc.key))
		if got != tc.want {
			t.Errorf("Get(%q) = %v, want %v", tc.key, got, tc.want)
		}
	}
}
