// Package trie is a naive trie data structure
package trie

type Trie struct {
	next     *Trie // 8 bytes (on 64-bit systems)
	children *Trie // 8 bytes
	key      rune  // 4 bytes
	// value is non-zero if the current node ends a key. This space would just
	// be used for padding anyway.
	value int // 8 bytes
}

func New() *Trie {
	return &Trie{}
}

func (tr *Trie) Insert(key string, value int) {
	e := tr
Rune:
	for _, r := range key {
		for c := e.children; c != nil; c = c.next {
			if c.key == r {
				e = c
				continue Rune
			}
		}
		e.children = &Trie{
			next: e.children,
			key:  r,
		}
		e = e.children
	}
	e.value = value
}

func (tr *Trie) Exists(key string) bool {
	return tr.Get(key) != 0
}

func (tr *Trie) Get(key string) int {
	e := tr
Rune:
	for _, r := range key {
		for c := e.children; c != nil; c = c.next {
			if c.key == r {
				e = c
				continue Rune
			}
		}
		return 0
	}
	return e.value
}
