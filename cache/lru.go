package cache

// LRU represents a cache that discards the least recently used items first.
// Items are stored in order of access in a doubly linked list.
// A previously unseen key is inserted at the front of the list.
// A duplicate key is deleted from the list and reinsert at the beginning.
// Remove operation deletes an element from the end and from the symbol table.
type LRU struct {
	first *lrunode
	last  *lrunode
	// st maps cache key to its location in linked list.
	st map[string]*lrunode
}
type lrunode struct {
	key   string
	value []byte
	prev  *lrunode
	next  *lrunode
}

// Set puts a key into the cache.
func (c *LRU) Set(key string, value []byte) {
	newnode := lrunode{
		key:   key,
		value: value,
	}

	n, ok := c.st[key]
	if ok {
		c.delete(n)
	}
	c.frontInsert(&newnode)
}

// delete removes a node from the linked list.
func (c *LRU) delete(n *lrunode) {
	delete(c.st, n.key)
	// Delete middle node.
	if c.first != n && c.last != n {
		n.prev.next = n.next
		n.next.prev = n.prev
		return
	}

	// Delete first node.
	if c.first == n {
		c.first = n.next
		if n.next != nil {
			n.next.prev = nil
		}
	}

	// Delete last node.
	if c.last == n {
		c.last = n.prev
		if n.prev != nil {
			n.prev.next = nil
		}
	}
}

// frontInsert adds a new node at the beginning of the linked list.
func (c *LRU) frontInsert(newnode *lrunode) {
	if c.first != nil {
		newnode.next = c.first
		c.first.prev = newnode
	}
	c.first = newnode
	if c.st == nil {
		c.st = make(map[string]*lrunode)
		c.last = newnode
	}
	c.st[newnode.key] = newnode
}

// Remove deletes and returns the least recently accessed key-value pair.
func (c *LRU) Remove() (key string, value []byte) {
	if c.last == nil {
		return
	}
	key, value = c.last.key, c.last.value
	c.delete(c.last)
	return
}
