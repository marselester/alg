package strsearch

// ASCIIRadix is extended ASCII alphabet size (number of characters).
const ASCIIRadix = 256

/*
Trie is a search tree known as a trie (re"trie"val), a data structure built from the characters of the string keys.
It's composed of nodes that contain links that are either null or references to other nodes.
Each node has R links (R is the alphabet size), often tries have a substantial number of null links.
Each link corresponds to a character value.
Each node has a corresponding value, which may be null or the value associated with one of the string keys in the symbol table.
Specifically, the value is stored in the node corresponding to its last character.

Nodes with null values exist to facilitate search in the trie and do not correspond to keys.

The shape of a trie is independent of the key insertion/deletion order.
Tries are optimal for search hit (get/put) — search time is proportional to the length of the search key.
Search miss doesn't depend on the key length and typically require examining just a few nodes.
The average number of nodes examined for search miss in a trie built from n random keys over an alphabet
of size R is ~log n (logarithm's base is R). For example, a search miss in a trie built from 1M random keys
will require examining only three or four nodes.

The number of links in a trie is between n * R and w * n * R where w is the average key length:

	when keys are short, the number of links is close to n * R
	when keys are long, the number of links is close to w * n * R
	decreasing R can save a huge amount of space

This trie implementation is not practical for large number of long keys taken from large alphabets,
because it will require space proportional to R times the total number of key characters.

The primary reason trie space is excessive for long keys because they tend to have long tails in the trie,
with each node having a single link to the next node (situation known as external one-way branching).
*/
type Trie struct {
	radix int
	root  *node
}
type node struct {
	value string
	next  []*node
}

// NewTrie creates an R-way trie (trie for an R-character alphabet).
func NewTrie(radix int) *Trie {
	return &Trie{
		radix: radix,
		root: &node{
			next: make([]*node, radix),
		},
	}
}

/*
Get searches the value associated with a given key.
Each node in the trie has a link corresponding to each possible string character.
We start at the root, then follow the link associated with the first character in the key;
from that node we follow the link associated with the second character in the key and so forth,
until reaching the last character of the key or a null link. At this point, one of three following
conditions holds:

	search hit (node value is not empty)
	search miss (node value is empty)
	search miss (search terminated with a null link)

*/
func (t *Trie) Get(key string) string {
	if key == "" {
		return ""
	}

	if n := t.get(t.root, key, 0); n != nil {
		return n.value
	}
	return ""
}

// get returns a node associated with key in the subtrie rooted at n.
func (t *Trie) get(n *node, key string, i int) *node {
	if n == nil {
		return nil
	}

	if isLastChar := i == len(key); isLastChar {
		return n
	}

	c := key[i]
	return t.get(n.next[c], key, i+1)
}

/*
Put inserts a key with value into a trie starting from a search from the first key character
until reaching the last character or a null link. At this point, one of the following two conditions holds:

	create a node (null link found before reaching the last character of the key)
	update node's value (found the last character of the key before reaching a null link)

*/
func (t *Trie) Put(key, value string) {
	if key == "" {
		return
	}

	n := t.root
	for i := 0; i < len(key); i++ {
		c := key[i]
		if n.next[c] == nil {
			n.next[c] = &node{
				next: make([]*node, t.radix),
			}
		}
		n = n.next[c]
	}
	n.value = value
}
