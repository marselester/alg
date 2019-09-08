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
