// Package trie implements a de la Briandais trie.
package trie

// Trie is a de la Briandais trie.
type Trie struct {
	root *node
}

// Each node requires more than 16 bytes (2 pointers plus a character).  In
// practice it appears the Go allocator is going to hand us 32 bytes even if we
// only needed 24. So we can have some flexibility up to 32 bytes.
type node struct {
	next      *node // 8 bytes (on 64-bit systems)
	children  *node // 8 bytes
	character rune  // 4 bytes (plus 4 bytes alignment padding)
	value     int   // 8 bytes
}

// New constructs an empty Trie.
func New() *Trie {
	return &Trie{root: &node{}}
}

// Insert stores a key in the Trie.
func (tr *Trie) Insert(key string, value int) {
	e := tr.root
Rune:
	for _, r := range key {
		for c := e.children; c != nil; c = c.next {
			if c.character == r {
				e = c
				continue Rune
			}
		}
		e.children = &node{
			next:      e.children,
			character: r,
		}
		e = e.children
	}
	e.value = value
}

// Exists returns whether a key exists in the Trie.
func (tr *Trie) Exists(key string) bool {
	return tr.Get(key) != 0
}

// Get returns the value stored at the key.
func (tr *Trie) Get(key string) int {
	e := tr.root
Rune:
	for _, r := range key {
		for c := e.children; c != nil; c = c.next {
			if c.character == r {
				e = c
				continue Rune
			}
		}
		return 0
	}
	return e.value
}
