package symboltable

import (
	"fmt"
	"testing"
)

func TestBinarySearchRank(t *testing.T) {
	var tt = []struct {
		st   BinarySearch
		key  string
		want int
	}{
		{
			st:   BinarySearch{},
			key:  "A",
			want: 0,
		},
		{
			st:   BinarySearch{keys: []string{"A"}},
			key:  "A",
			want: 0,
		},
		{
			st:   BinarySearch{keys: []string{"A", "B"}},
			key:  "C",
			want: 2,
		},
		{
			st:   BinarySearch{keys: []string{"A", "D"}},
			key:  "C",
			want: 1,
		},
		{
			st:   BinarySearch{keys: []string{"A", "B"}},
			key:  "B",
			want: 1,
		},
		{
			st:   BinarySearch{keys: []string{"A", "B"}},
			key:  "A",
			want: 0,
		},
		{
			st:   BinarySearch{keys: []string{"A", "C", "E", "H", "L", "M", "P", "R", "S", "X"}},
			key:  "P",
			want: 6,
		},
		{
			st:   BinarySearch{keys: []string{"A", "C", "E", "H", "L", "M", "P", "R", "S", "X"}},
			key:  "Q",
			want: 7,
		},
	}
	for _, tc := range tt {
		got := tc.st.rank(tc.key)
		if got != tc.want {
			t.Errorf("rank(%q) = %d, want %d", tc.key, got, tc.want)
		}
	}
}

func TestBinarySearchGet(t *testing.T) {
	tt := []struct {
		st        BinarySearch
		searchKey string
		want      int
	}{
		{
			st:        BinarySearch{},
			searchKey: "A",
			want:      -1,
		},
		{
			st: BinarySearch{
				keys:   []string{"A"},
				values: []int{1},
			},
			searchKey: "A",
			want:      1,
		},
		{
			st: BinarySearch{
				keys:   []string{"A", "B"},
				values: []int{1, 2},
			},
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

func TestBinarySearchPut(t *testing.T) {
	tt := []struct {
		st       BinarySearch
		key      string
		value    int
		wantKeys []string
		wantVals []int
	}{
		{
			st:       BinarySearch{},
			key:      "A",
			value:    123,
			wantKeys: []string{"A"},
			wantVals: []int{123},
		},
		{
			st: BinarySearch{
				keys:   []string{"B"},
				values: []int{456},
			},
			key:      "A",
			value:    123,
			wantKeys: []string{"A", "B"},
			wantVals: []int{123, 456},
		},
		{
			st: BinarySearch{
				keys:   []string{"A"},
				values: []int{123},
			},
			key:      "B",
			value:    456,
			wantKeys: []string{"A", "B"},
			wantVals: []int{123, 456},
		},
	}
	for _, tc := range tt {
		tc.st.Put(tc.key, tc.value)
		if !equal(tc.st.keys, tc.wantKeys) {
			t.Errorf("Put(%q, %d) got keys %v, want %v", tc.key, tc.value, tc.st.keys, tc.wantKeys)
		}
		if fmt.Sprint(tc.st.values) != fmt.Sprint(tc.wantVals) {
			t.Errorf("Put(%q, %d) got vals %v, want %v", tc.key, tc.value, tc.st.values, tc.wantVals)
		}
	}
}

func TestBinarySearchAfterPuts(t *testing.T) {
	st := BinarySearch{}
	for v, k := range "SEARCHEXAMPLE" {
		st.Put(string(k), v)
	}

	wantKeys := []string{"A", "C", "E", "H", "L", "M", "P", "R", "S", "X"}
	if !equal(st.keys, wantKeys) {
		t.Errorf("Put() got keys %v, want %v", st.keys, wantKeys)
	}

	wantVals := []int{8, 4, 12, 5, 11, 9, 10, 3, 0, 7}
	if fmt.Sprint(st.values) != fmt.Sprint(wantVals) {
		t.Errorf("Put() got vals %v, want %v", st.values, wantVals)
	}
}
