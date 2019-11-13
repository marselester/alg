package strsearch_test

import (
	"fmt"

	"github.com/marselester/alg/strsearch"
)

func ExampleTrie() {
	st := strsearch.NewTrie(strsearch.ASCIIRadix)
	st.Put("fizz", "bazz")
	fmt.Print(st.Get("fizz"))
	// Output: bazz
}

func ExampleTernaryTrie() {
	tst := strsearch.TernaryTrie{}
	keys := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}
	for i, k := range keys {
		tst.Put(k, fmt.Sprintf("%d", i))
	}

	for _, k := range keys {
		fmt.Printf("%s %s\n", k, tst.Get(k))
	}
	// Output:
	// she 0
	// sells 1
	// sea 6
	// shells 3
	// by 4
	// the 5
	// sea 6
	// shore 7
}
