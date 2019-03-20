package hashtable

import symboltable "github.com/marselester/alg/search/symbol-table"

// DefaultChainingSize is a default length of the hash table's underlying array.
const DefaultChainingSize = 97

// Hasher function should return a hash value for given key so it fits into [0; size-1] range.
type Hasher func(key string, size int) int

type configOption func(*SeparateChaining)

// WithChainingSize defines length of the hash table's underlying array.
func WithChainingSize(size int) configOption {
	return func(ht *SeparateChaining) {
		ht.size = size
	}
}

// WithHash defines a hash function to be used in the hash table.
func WithHash(hash Hasher) configOption {
	return func(ht *SeparateChaining) {
		ht.hash = hash
	}
}

// NewSeparateChaining returns a hashing with separate chaining.
func NewSeparateChaining(options ...configOption) *SeparateChaining {
	var ht SeparateChaining
	for _, opt := range options {
		opt(&ht)
	}
	if ht.size <= 0 {
		ht.size = DefaultChainingSize
	}
	if ht.hash == nil {
		ht.hash = Hash
	}
	ht.a = make([]symboltable.SequentialSearch, ht.size)
	return &ht
}
