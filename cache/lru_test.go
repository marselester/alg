package cache

import (
	"fmt"
	"testing"
)

// cachedLRUKeyVal returns a list of key-value pairs from LRU linked list for testing.
func cachedLRUKeyVal(c *LRU) []string {
	var kv []string
	for n := c.first; n != nil; n = n.next {
		kv = append(kv, fmt.Sprintf("%s:%s", n.key, n.value))
	}
	return kv
}

func TestLRUSet(t *testing.T) {
	var tt = []struct {
		cache LRU
		items []struct {
			key   string
			value []byte
		}
		want []string
	}{
		{
			cache: LRU{},
			items: []struct {
				key   string
				value []byte
			}{
				{"a", []byte("A")},
			},
			want: []string{"a:A"},
		},
		{
			cache: LRU{},
			items: []struct {
				key   string
				value []byte
			}{
				{"b", []byte("B")},
				{"a", []byte("A")},
			},
			want: []string{"a:A", "b:B"},
		},
		{
			cache: LRU{},
			items: []struct {
				key   string
				value []byte
			}{
				{"a", []byte("A")},
				{"a", []byte("AAA")},
			},
			want: []string{"a:AAA"},
		},
		{
			cache: LRU{},
			items: []struct {
				key   string
				value []byte
			}{
				{"b", []byte("B")},
				{"a", []byte("A")},
				{"a", []byte("AAA")},
			},
			want: []string{"a:AAA", "b:B"},
		},
		{
			cache: LRU{},
			items: []struct {
				key   string
				value []byte
			}{
				{"a", []byte("A")},
				{"b", []byte("B")},
				{"a", []byte("AAA")},
			},
			want: []string{"a:AAA", "b:B"},
		},
	}

	for _, tc := range tt {
		for _, item := range tc.items {
			tc.cache.Set(item.key, item.value)
		}
		got := cachedLRUKeyVal(&tc.cache)
		if !equal(got, tc.want) {
			t.Errorf("Set() = %v, want %v", got, tc.want)
		}
	}
}

func TestLRURemove(t *testing.T) {
	var tt = []struct {
		cache LRU
		items []struct {
			key   string
			value []byte
		}
		want string
	}{
		{
			cache: LRU{},
			items: nil,
			want:  ":",
		},
		{
			cache: LRU{},
			items: []struct {
				key   string
				value []byte
			}{
				{"a", []byte("A")},
			},
			want: "a:A",
		},
		{
			cache: LRU{},
			items: []struct {
				key   string
				value []byte
			}{
				{"b", []byte("B")},
				{"a", []byte("A")},
			},
			want: "b:B",
		},
		{
			cache: LRU{},
			items: []struct {
				key   string
				value []byte
			}{
				{"a", []byte("A")},
				{"b", []byte("B")},
				{"c", []byte("C")},
			},
			want: "a:A",
		},
	}

	for _, tc := range tt {
		for _, item := range tc.items {
			tc.cache.Set(item.key, item.value)
		}
		k, v := tc.cache.Remove()
		got := fmt.Sprintf("%s:%s", k, v)
		if got != tc.want {
			t.Errorf("Remove() = %s, want %v", got, tc.want)
		}
	}
}
