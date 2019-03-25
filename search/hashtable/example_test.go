package hashtable_test

import (
	"fmt"

	"github.com/marselester/alg/search/hashtable"
)

func Example() {
	ht := hashtable.NewSeparateChaining(
		hashtable.WithTableSize(997),
		hashtable.WithHash(hashtable.Hash),
	)
	ht.Put("age", 100)
	fmt.Printf("%d\n", ht.Get("name"))
	fmt.Printf("%d\n", ht.Get("age"))
	// Output:
	// -1
	// 100
}
