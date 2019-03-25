/*
Package hashtable implements a symbol table using an array.
A hash function transforms the search key into an array index.
Different keys may hash to the same array index.
A collision-resolution process (separate chaining, linear probing) deals with this situation.
With hashing, search and insert operations require constant (amortized) time.

Suppose an array can hold m key-value pairs, then a hash function should transform
any key into an array index in the range [0; m-1]. It should uniformly distribute the keys.
Modular hashing is the most commonly used method for hashing integers:
choose array length m to be prime, for any positive integer k compute remainder k % m.
If m is not prime, not all bits of the key play a role, missing opportunity to disperse the values evenly.
*/
package hashtable

import symboltable "github.com/marselester/alg/search/symbol-table"

// Hash computes a modular hash function for a key using Horner's method.
// Note, size should be a prime integer.
// Mersenne prime (e.g., 31, 127 or 8191) is a prime number
// that is one less than a power of 2. This means that the mod can be done with one shift
// and one subtract if the machine's multiply instruction is slow
// https://algs4.cs.princeton.edu/34hash/.
func Hash(key string, size int) int {
	h := 0
	for _, r := range key {
		// A sufficiently small (to prevent overflow) prime integer 31 ensures
		// that the bits of all the runes play a role. Bitwise "and" operation
		// h & (size-1) can be used instead of "%" remainder operation.
		h = (h*31 + int(r)) % size
	}
	return h
}

// SeparateChaining is a hash table implementation based on separate chaining
// collision-resolution process: build a linked-list of key-value pairs
// for each of the array indices. Items that collide are chained together.
// Since we have m lists and n keys, the average length of the lists is always n/m.
// The number of compares for search miss and insert is ~n/m.
// Array resizing can be used to make sure the lists are short no matter how many keys are stored.
type SeparateChaining struct {
	a []symboltable.SequentialSearch
	config
	// size is a length of the hash table's underlying array.
	// size int
	// hash Hasher
}

// Put uses a hash function to choose a list for the key.
// Then it scans through the list and updates the value associated with the search key.
// If key is not found, a new node is inserted at the beginning of the list.
func (ht *SeparateChaining) Put(key string, value int) {
	index := ht.hash(key, ht.size)
	ht.a[index].Put(key, value)
}

// Get uses a hash function to choose a list for the key.
// Then it scans through the list and compares the search key with key in each node.
func (ht *SeparateChaining) Get(key string) int {
	index := ht.hash(key, ht.size)
	return ht.a[index].Get(key)
}

// LinearProbing is a hash table implementation based on linear probing
// collision-resolution process: when there is a collision, check the next entry in the table
// (increment the index, wrap back to the beginning of the table if the end is reached)
// until finding either the search key or an empty entry.
//
// α = n/m is a percentage of table entries that are occupied (load factor).
// It must not reach 1 (full table).
// For good performance the load factor should be between 1/8 and 1/2 by using array resizing.
type LinearProbing struct {
	keys   []string
	values []int
	// n is the number of key-value pairs in the table.
	n int
	config
}

// Put stores a key in the hash table:
// if a new key hashes to an empty entry (blank string), it's stored there;
// if not, it scans sequentially to find an empty position.
func (ht *LinearProbing) Put(key string, value int) {
	// Double the size of linear-probing table.
	if ht.n >= ht.size/2 {
		ht.resize(ht.size * 2)
	}

	i := ht.hash(key, ht.size)
	for ; ht.keys[i] != ""; i = (i + 1) % ht.size {
		if key == ht.keys[i] {
			ht.values[i] = value
			return
		}
	}

	ht.keys[i] = key
	ht.values[i] = value
	ht.n++
}

// resize puts old keys from the old table into the new one
// by rehashing all the keys.
func (ht *LinearProbing) resize(size int) {
	newht := NewLinearProbing(
		WithTableSize(size),
		WithHash(ht.hash),
	)
	for i := 0; i < ht.size; i++ {
		if ht.keys[i] != "" {
			newht.Put(ht.keys[i], ht.values[i])
		}
	}
	ht.keys = newht.keys
	ht.values = newht.values
	ht.size = newht.size
}

// Get searches for a key sequentially starting at its hash index
// until finding an empty string (search miss) or the key (search hit).
func (ht *LinearProbing) Get(key string) int {
	i := ht.hash(key, ht.size)
	for ; ht.keys[i] != ""; i = (i + 1) % ht.size {
		if key == ht.keys[i] {
			return ht.values[i]
		}
	}
	return -1
}
