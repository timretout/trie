// Package trie is a naive trie data structure
package trie

const end rune = -1

type Trie struct {
	root *element
}

type element struct {
	next     *element
	children *element
	v        rune
}

func New() *Trie {
	return &Trie{
		root: &element{},
	}
}

func (tr *Trie) Insert(s string) {
	e := tr.root
Rune:
	for _, r := range s {
		for c := e.children; c != nil; c = c.next {
			if c.v == r {
				e = c
				continue Rune
			}
		}
		e.children = &element{
			next: e.children,
			v:    r,
		}
		e = e.children
	}
	for c := e.children; c != nil; c = c.next {
		if c.v == end {
			return
		}
	}
	e.children = &element{
		next: e.children,
		v:    end,
	}
}

func (tr *Trie) Exists(s string) bool {
	e := tr.root
Rune:
	for _, r := range s {
		for c := e.children; c != nil; c = c.next {
			if c.v == r {
				e = c
				continue Rune
			}
		}
		return false
	}
	for c := e.children; c != nil; c = c.next {
		if c.v == end {
			return true
		}
	}
	return false
}
