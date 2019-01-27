package symboltable

import (
	"fmt"
	"testing"
)

func TestSequentialSearchGet(t *testing.T) {
	tt := []struct {
		st        SequentialSearch
		searchKey string
		want      int
	}{
		{
			st:        SequentialSearch{},
			searchKey: "A",
			want:      -1,
		},
		{
			st: SequentialSearch{first: &node{
				key:   "A",
				value: 1,
			}},
			searchKey: "A",
			want:      1,
		},
		{
			st: SequentialSearch{first: &node{
				key:   "A",
				value: 1,
				next: &node{
					key:   "B",
					value: 2,
				},
			}},
			searchKey: "B",
			want:      2,
		},
	}
	for _, tc := range tt {
		got := tc.st.Get(tc.searchKey)
		if got != tc.want {
			t.Errorf("Get(%q) = %d, want %d", tc.searchKey, got, tc.want)
		}
	}
}

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

// items returns a slice of key:value pairs from a linked list.
func items(n *node) []string {
	var kv []string
	for ; n != nil; n = n.next {
		kv = append(kv, fmt.Sprintf("%s:%d", n.key, n.value))
	}
	return kv
}

func TestSequentialSearchPut(t *testing.T) {
	tt := []struct {
		st    SequentialSearch
		key   string
		value int
		want  []string
	}{
		{
			st:    SequentialSearch{},
			key:   "A",
			value: 1,
			want:  []string{"A:1"},
		},
		{
			st: SequentialSearch{first: &node{
				key:   "A",
				value: 1,
			}},
			key:   "A",
			value: 2,
			want:  []string{"A:2"},
		},
		{
			st: SequentialSearch{first: &node{
				key:   "A",
				value: 1,
			}},
			key:   "B",
			value: 2,
			want:  []string{"B:2", "A:1"},
		},
	}
	for _, tc := range tt {
		tc.st.Put(tc.key, tc.value)
		got := items(tc.st.first)
		if !equal(got, tc.want) {
			t.Errorf("Put(%q, %d) got %v, want %v", tc.key, tc.value, got, tc.want)
		}
	}
}
