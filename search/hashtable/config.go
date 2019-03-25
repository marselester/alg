package hashtable

import symboltable "github.com/marselester/alg/search/symbol-table"

// DefaultTableSize is a default length of the hash table's underlying array.
const DefaultTableSize = 97

// Hasher function should return a hash value for given key so it fits into [0; size-1] range.
type Hasher func(key string, size int) int

type config struct {
	// size is a length of the hash table's underlying array.
	size int
	hash Hasher
}
type configOption func(*config)

// WithTableSize defines length of the hash table's underlying array.
func WithTableSize(size int) configOption {
	return func(c *config) {
		c.size = size
	}
}

// WithHash defines a hash function to be used in the hash table.
func WithHash(hash Hasher) configOption {
	return func(c *config) {
		c.hash = hash
	}
}

// NewSeparateChaining returns a hashing with separate chaining.
func NewSeparateChaining(options ...configOption) *SeparateChaining {
	var ht SeparateChaining
	for _, opt := range options {
		opt(&ht.config)
	}
	if ht.size <= 0 {
		ht.size = DefaultTableSize
	}
	if ht.hash == nil {
		ht.hash = Hash
	}
	ht.a = make([]symboltable.SequentialSearch, ht.size)
	return &ht
}

// NewLinearProbing returns a hashing with linear probing (open addressing).
func NewLinearProbing(options ...configOption) *LinearProbing {
	var ht LinearProbing
	for _, opt := range options {
		opt(&ht.config)
	}
	if ht.size <= 0 {
		ht.size = DefaultTableSize
	}
	if ht.hash == nil {
		ht.hash = Hash
	}
	ht.keys = make([]string, ht.size)
	ht.values = make([]int, ht.size)
	return &ht
}
