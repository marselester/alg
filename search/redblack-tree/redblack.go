/*
Package redblack implements a red-black binary search tree (BST) using recursive approach which is limited to a stack size.
BST is a binary tree where key in a node is larger than the keys in all its left children and
smaller than the keys in right children. New nodes are attached at the bottom of the tree.

Red-black tree encodes 2-3 tree with standard BST (2-nodes).
A 2-3 tree contains 2-nodes (1 key, 2 links) or 3-nodes (2 keys, 3 links) whose
left link has smaller keys, a middle link has keys in between, a right link has larger keys.

A red link binds together two standard nodes to represent a 3-node.
A black link binds together the 2-3 tree.
Red links lean left.
No node has two red links connected to it.
Every path from the root has the same number of black links (black height) — perfect black balance.

Such representation allows to use Get method from standard BST (lg n) and
the efficient (lg n) insertion-balancing method from 2-3 trees.

This implementation allows right-leaning red links or two red links in a row during an operation,
but it always corrects these conditions before completion using rotation.

After insertion perform balancing operations on the way up the tree:

	rotate left if the right child is red and the left child is black
	rotate right if both the left child and its left child are red
	flip colors if both children are red
*/
package redblack

const (
	red   = true
	black = false
)

// Tree represents a red-black binary search tree.
type Tree struct {
	root *node
}
type node struct {
	// key is a unique comparable key, e.g., name.
	key string
	// value is a value associated with the key, e.g., Bob.
	value []byte
	// color of the link from parent to this node (red or black).
	color bool
	// left is pointer to the left subtree where smaller keys are stored.
	left *node
	// right is pointer to the right subtree where larger keys are stored.
	right *node
}

// isRed returns true if its link to parent is red.
func (n *node) isRed() bool {
	if n == nil {
		return false
	}
	return n.color == red
}

// Get retrieves a key from the tree.
func (t *Tree) Get(key string) []byte {
	found := search(key, t.root)
	if found == nil {
		return nil
	}
	return found.value
}

// Set stores the key in the tree. First it looks up the key and if found, updates the value.
// If the key is new, it will be added to the tree.
// The root is colored black after each insertion: a red root implies that the root is part of a 3-node,
// but that's not the case.
func (t *Tree) Set(key string, value []byte) {
	t.root = put(key, value, t.root)
	t.root.color = black
}

// Keys returns all keys sorted in ascending order.
func (t *Tree) Keys() []string {
	return keys(nil, t.root)
}

// search recursively looks up node by key starting from node n.
func search(key string, n *node) *node {
	switch {
	// Search miss.
	case n == nil:
		return nil
	// Search hit.
	case key == n.key:
		return n
	// Check smaller keys on the left side.
	case key < n.key:
		return search(key, n.left)
	// Check larger keys on the right side.
	case key > n.key:
		return search(key, n.right)
	}
	return nil
}

// put updates the value of found node which was looked up by key.
// If key is not found, the new node with red link is added to the tree.
func put(key string, value []byte, n *node) *node {
	if n == nil {
		return &node{key: key, value: value, color: red}
	}

	if key < n.key {
		n.left = put(key, value, n.left)
	} else if key > n.key {
		n.right = put(key, value, n.right)
	} else {
		n.value = value
	}

	// Balance the tree on the way up the search path.
	if n.right.isRed() && !n.left.isRed() {
		n = rotateLeft(n)
	}
	if n.left.isRed() && n.left.left.isRed() {
		n = rotateRight(n)
	}
	if n.left.isRed() && n.right.isRed() {
		flipColors(n)
	}
	return n
}

// rotateLeft takes a node whose right link is red and returns a node with the same keys
// whose left link is red. For example, there is a right-leaning red link that needs to be
// rotated to lean to the left. Smaller key H (parent) becomes a child of larger key X.
// Before: H has a right child X. After: X has a left child H.
func rotateLeft(h *node) *node {
	// Left (less than H) and rightmost (greater than X) nodes stay untouched.
	// Middle node (between H and X) changes its parent X to H.
	x := h.right
	h.right = x.left
	// Parent H becomes a left child of its right child X.
	x.left = h
	// After rotation X gets color of H (could be red or black).
	x.color = h.color
	// H is now leaning left, so the link is marked red.
	h.color = red
	return x
}

// rotateRight converts a left-leaning red link to a right-leaning one.
// Larger key H (parent) becomes a child of smaller key X.
// Before: X has a left child H. After: H has a right child X.
func rotateRight(h *node) *node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	return x
}

// flipColors flips the colors of two red children of node h to black.
// Parent's color is also flipped from black to red.
func flipColors(h *node) {
	h.color = red
	h.left.color = black
	h.right.color = black
}

// keys recursively traverses the tree and returns all keys in order.
// Get all left children (smaller). Get the current key. Get all right children (larger).
func keys(kk []string, n *node) []string {
	if n == nil {
		return kk
	}
	kk = keys(kk, n.left)
	kk = append(kk, n.key)
	kk = keys(kk, n.right)
	return kk
}
