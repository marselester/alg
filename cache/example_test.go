package cache_test

import (
	"fmt"

	"github.com/marselester/alg/cache"
)

func Example() {
	c := cache.LRU{}
	c.Set("a", []byte("A"))
	c.Set("b", []byte("B"))
	c.Set("c", []byte("C"))
	c.Set("d", []byte("D"))
	c.Set("a", []byte("AA"))
	c.Set("b", []byte("BB"))
	c.Set("e", []byte("E"))
	c.Set("a", []byte("AAA"))
	c.Set("b", []byte("BBB"))
	c.Set("c", []byte("CC"))
	c.Set("d", []byte("DD"))
	c.Set("e", []byte("EE"))
	for {
		k, v := c.Remove()
		if k == "" {
			break
		}
		fmt.Printf("%s:%s\n", k, v)
	}
	// Output:
	// a:AAA
	// b:BBB
	// c:CC
	// d:DD
	// e:EE
}
