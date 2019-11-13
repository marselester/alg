package strsearch

import (
	"unicode/utf8"
)

/*
TernaryTrie is a ternary search trie (TST) which helps to avoid the excessive space cost associated with R-way tries.
In a TST, each node has a character, three links, and a value.
The three links correspond to keys whose current characters are less than, equal to, or greater than the node's character.

Using this arrangement is equivalent to implementing R-way trie node as a binary search tree.
BST representation of each trie node depends on the order of key insertion.

TST requires far less space than R-way trie: the number of links in a TST build from n string keys
of average length w is between 3*n and 3*w*n.

A search miss in a TST build from n random string keys on average requires ln n character compares.
A search hit or an insertion in a TST uses ln n + L character compares, where L is the length of the search key.
*/
type TernaryTrie struct {
	root *tstnode
}
type tstnode struct {
	char rune
	// value associated with a key.
	value string
	left  *tstnode
	mid   *tstnode
	right *tstnode
}

/*
Get searches the value associated with a given key.
To search, the first character in the key is compared with the character at the root:

	if it's less, take the left link
	if it's greater, take the right link
	if it's equal, take the middle link and move to the next search key charater

Search hit: the node where the search ends has non-blank value.
Search miss: a null link is encountered or if the node where the search ends has a blank value.
*/
func (t *TernaryTrie) Get(key string) string {
	if key == "" {
		return ""
	}

	n := t.get(t.root, key, 0)
	if n == nil {
		return ""
	}
	return n.value
}

func (t *TernaryTrie) get(n *tstnode, key string, i int) *tstnode {
	if n == nil {
		return nil
	}

	char, width := utf8.DecodeRuneInString(key[i:])
	switch {
	case char < n.char:
		return t.get(n.left, key, i)
	case char > n.char:
		return t.get(n.right, key, i)
	}

	if isLastChar := i+width == len(key); isLastChar {
		return n
	}
	return t.get(n.mid, key, i+width)
}

/*
Put inserts a key with value into a trie starting from a search from the first key character
until reaching the last character or a null link. At this point, one of the following two conditions holds:

	create a node (null link found before reaching the last character of the key)
	update node's value (found the last character of the key before reaching a null link)

*/
func (t *TernaryTrie) Put(key, value string) {
	t.root = t.put(t.root, key, value, 0)
}

func (t *TernaryTrie) put(n *tstnode, key, value string, i int) *tstnode {
	char, width := utf8.DecodeRuneInString(key[i:])

	if n == nil {
		n = &tstnode{char: char}
	}

	switch {
	case char < n.char:
		n.left = t.put(n.left, key, value, i)
	case char > n.char:
		n.right = t.put(n.right, key, value, i)
	case char == n.char:
		if isLastChar := i+width == len(key); isLastChar {
			n.value = value
		} else {
			n.mid = t.put(n.mid, key, value, i+width)
		}
	}

	return n
}
