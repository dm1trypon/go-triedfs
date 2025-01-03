// Package triedfs implements a Trie data structure
// with depth-first search capabilities for adding
// and searching for sequences of comparable values.
package triedfs

// Trie represents a trie structure, holding a root node.
type Trie[V comparable] struct {
	root *node[V]
}

// NewTrie initializes and returns a new Trie instance.
func NewTrie[V comparable]() *Trie[V] {
	return &Trie[V]{
		root: newNode[V](),
	}
}

// Add inserts a sequence of values into the trie.
func (t *Trie[V]) Add(values []V) {
	t.root.add(values, 0)
}

// Search checks if the given sequence of values exists in the trie.
// Returns true if found, otherwise false.
func (t *Trie[V]) Search(values []V) bool {
	if len(values) == 0 {
		return false // No empty sequences are considered found.
	}
	return t.root.search(values, 0)
}

// node represents a single node in the trie,
// storing children nodes mapped by their corresponding value.
type node[V comparable] struct {
	nodes map[V]*node[V]
	isEnd bool
}

// newNode creates and initializes a new node with an empty children map.
func newNode[V comparable]() *node[V] {
	return &node[V]{
		nodes: make(map[V]*node[V]),
	}
}

// add recursively inserts a sequence of values into the node.
func (n *node[V]) add(val []V, idx int) {
	// Base case: if the end of the sequence is reached, nothing more to add.
	if idx == len(val) {
		n.isEnd = true
		return
	}

	// Create or get the next node for the current value.
	if n.nodes[val[idx]] == nil {
		n.nodes[val[idx]] = newNode[V]()
	}
	next := n.nodes[val[idx]] // Move to the child node.

	idx++              // Proceed to the next index.
	next.add(val, idx) // Recursively add the remaining values.
}

// search recursively checks if a sequence of values exists in the node.
func (n *node[V]) search(val []V, idx int) bool {
	// Base case: if the end of the sequence is reached, the sequence is found.
	if idx == len(val) {
		return n.isEnd
	}

	// If there is no corresponding child node for the current value, return false.
	if n.nodes[val[idx]] == nil {
		return false
	}
	next := n.nodes[val[idx]] // Move to the next node.

	idx++                        // Proceed to the next index.
	return next.search(val, idx) // Recursively search for the remaining values.
}
