package redblack_test

import (
	"fmt"

	"github.com/marselester/alg/search/redblack-tree"
)

func Example() {
	tree := redblack.Tree{}
	tree.Set("name", []byte("Bob"))
	fmt.Printf("%s", tree.Get("name"))
	// Output: Bob
}
