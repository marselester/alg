// Package cache implements a cache using move-to-front strategy,
// where items that have been recently accessed are more likely to be reaccessed.
package cache

// MoveToFront represents a cache that stores keys using move-to-front strategy.
// A previously unseen key is inserted at the front of the list.
// A duplicate key is deleted from the list and reinsert at the beginning.
type MoveToFront struct {
	first *node
}

type node struct {
	key   string
	value []byte
	next  *node
}

// Set puts a key into the cache.
func (c *MoveToFront) Set(key string, value []byte) {
	newnode := node{
		key:   key,
		value: value,
	}
	c.deleteDuplicate(key)
	c.frontInsert(&newnode)
}

// deleteDuplicate deletes a node by given key found in the linked list.
func (c *MoveToFront) deleteDuplicate(key string) {
	var prev *node
	for n := c.first; n != nil; {
		if n.key == key {
			if prev == nil {
				c.first = n.next // Delete first node.
			} else {
				prev.next = n.next // Delete middle node.
			}
			break
		}
		prev = n
		n = n.next
	}
}

// frontInsert adds a new node at the beginning of the linked list.
func (c *MoveToFront) frontInsert(newnode *node) {
	newnode.next = c.first
	c.first = newnode
}

// Get retrieves a key from the cache.
func (c *MoveToFront) Get(key string) []byte {
	for n := c.first; n != nil; n = n.next {
		if n.key == key {
			return n.value
		}
	}
	return nil
}
