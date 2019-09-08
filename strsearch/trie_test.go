package strsearch

import "testing"

func TestTrieGet(t *testing.T) {
	st := NewTrie(ASCIIRadix)
	st.root = &node{
		value: "root",
		next:  make([]*node, st.radix),
	}
	st.root.next['a'] = &node{
		next: make([]*node, st.radix),
	}
	st.root.next['a'].next['b'] = &node{
		value: "ab",
		next:  make([]*node, st.radix),
	}

	tests := map[string]struct {
		input string
		want  string
	}{
		"ab key found":        {"ab", "ab"},
		"a key not found":     {"a", ""},
		"b key not found":     {"b", ""},
		"abc key not found":   {"abc", ""},
		"empty key not found": {"", ""},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := st.Get(tc.input)
			if got != tc.want {
				t.Errorf("Trie.Get(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestTriePut(t *testing.T) {
	st := NewTrie(ASCIIRadix)

	st.Put("ab", "fizz")
	want := ""
	if st.root.value != want {
		t.Errorf("Trie.Put(ab, fizz) root value is %q, want %q", st.root.value, want)
	}
	if st.root.next['a'].value != want {
		t.Errorf("Trie.Put(ab, fizz) 'a' value is %q, want %q", st.root.next['a'].value, want)
	}
	want = "fizz"
	if st.root.next['a'].next['b'].value != want {
		t.Errorf("Trie.Put(ab, fizz) 'ab' value is %q, want %q", st.root.next['a'].next['b'].value, want)
	}

	st.Put("ab", "bazz")
	want = ""
	if st.root.value != want {
		t.Errorf("Trie.Put(ab, bazz) root value is %q, want %q", st.root.value, want)
	}
	if st.root.next['a'].value != want {
		t.Errorf("Trie.Put(ab, bazz) 'a' value is %q, want %q", st.root.next['a'].value, want)
	}
	want = "bazz"
	if st.root.next['a'].next['b'].value != want {
		t.Errorf("Trie.Put(ab, bazz) 'ab' value is %q, want %q", st.root.next['a'].next['b'].value, want)
	}
}
