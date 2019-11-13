package strsearch

import (
	"testing"
)

func TestTernaryTrie_Get(t *testing.T) {
	tst := abcTST()

	tests := []struct {
		key       string
		wantValue string
	}{
		{"", ""},
		{"fizz", ""},
		{"sea", "path=shela key=sea"},
		{"she", "path=she key=she"},
		{"sells", "path=shells key=sells"},
		{"shells", "path=shells key=shells"},
	}

	for _, tc := range tests {
		got := tst.Get(tc.key)
		if got != tc.wantValue {
			t.Errorf("TernaryTrie.Get(%q) got value %q, want %q", tc.key, got, tc.wantValue)
		}
	}
}

func TestTernaryTrie_Put(t *testing.T) {
	tst := abcTST()

	tests := []struct {
		key       string
		value     string
		wantValue string
	}{
		{"", "", ""},
		{" ", "whitespace", "whitespace"},
		{"fizz", "", ""},
		{"", "fizz", ""},
		{"sea", "ocean", "ocean"},
		{"she", "👩", "👩"},
		{" 👩 ", "she", "she"},
	}

	for _, tc := range tests {
		tst.Put(tc.key, tc.wantValue)

		got := tst.Get(tc.key)
		if got != tc.wantValue {
			t.Errorf("TernaryTrie.Put(%q, %q) got value %q, want %q", tc.key, tc.value, got, tc.wantValue)
		}
	}
}

/*
abcTST returns the following ternary search trie:

	s
   /|\
  b h t
   /|\
  e e u
  | |
  l l
 /| |
a l l
  | |
  s s
*/
func abcTST() *TernaryTrie {
	root := tstnode{
		char:  's',
		value: "path=s",
		left: &tstnode{
			char:  'b',
			value: "path=sb",
		},
		mid: &tstnode{
			char:  'h',
			value: "path=sh",
			left: &tstnode{
				char:  'e',
				value: "path=she",
				mid: &tstnode{
					char:  'l',
					value: "path=shel",
					left: &tstnode{
						char:  'a',
						value: "path=shela key=sea",
					},
					mid: &tstnode{
						char:  'l',
						value: "path=shell",
						mid: &tstnode{
							char:  's',
							value: "path=shells key=sells",
						},
					},
				},
			},
			mid: &tstnode{
				char:  'e',
				value: "path=she key=she",
				mid: &tstnode{
					char:  'l',
					value: "path=shel",
					mid: &tstnode{
						char:  'l',
						value: "path=shell",
						mid: &tstnode{
							char:  's',
							value: "path=shells key=shells",
						},
					},
				},
			},
			right: &tstnode{
				char:  'u',
				value: "path=shu",
			},
		},
		right: &tstnode{
			char:  't',
			value: "path=st",
		},
	}
	return &TernaryTrie{
		root: &root,
	}
}
