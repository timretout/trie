// Package trie is a naive trie data structure
package trie

const end rune = -1

type Trie struct {
	root *element
}

type element struct {
	children []*element
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
		for _, c := range e.children {
			if c.v == r {
				e = c
				continue Rune
			}
		}
		c := element{v: r}
		e.children = append(e.children, &c)
		e = &c
	}
	term := element{v: end}
	e.children = append(e.children, &term)
}

func (tr *Trie) Exists(s string) bool {
	e := tr.root
Rune:
	for _, r := range s {
		for _, c := range e.children {
			if c.v == r {
				e = c
				continue Rune
			}
		}
		return false
	}
	for _, c := range e.children {
		if c.v == end {
			return true
		}
	}
	return false
}
