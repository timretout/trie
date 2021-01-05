// Package trie implements a pure trie
package trie

import (
	"strings"
)

type Trie struct {
	root *node
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alphabetLength = len(alphabet)

// A node is a vector of entries
type node [alphabetLength + 1]*entry // 53 * 8 bytes

// An entry is either a key or a link
type entry struct {
	key  string // 16 bytes + data
	link *node  // 8 byte pointer
}

func New() *Trie {
	return &Trie{root: &node{}}
}

func (tr *Trie) Insert(s string) {
	p := tr.root

	for i := 0; i <= len(s); i++ {
		var k int
		if i == len(s) {
			k = alphabetLength
		} else {
			k = strings.IndexByte(alphabet, s[i])
		}
		if k == -1 {
			panic("k is -1 for " + string(s[i]) + " while processing " + s)
		}

		x := p[k]
		if x == nil {
			p[k] = &entry{key: s}
			return
		} else if x.key == s {
			return
		} else if x.link == nil {
			x.link = &node{}
			var u int
			if i+1 >= len(x.key) {
				u = alphabetLength
			} else {
				u = strings.IndexByte(alphabet, x.key[i+1])
			}
			if u == -1 {
				panic("u is -1 for " + string(x.key[i+1]) + " while processing " + s)
			}

			x.link[u] = &entry{key: x.key}
			x.key = ""
		}
		p = x.link
	}
}

func (tr *Trie) Exists(s string) bool {
	p := tr.root

	for i := 0; i < len(s); i++ {
		k := strings.IndexByte(alphabet, s[i])

		x := p[k]
		if x == nil {
			return false
		} else if x.link != nil {
			p = x.link
		} else {
			return x.key == s
		}
	}

	x := p[alphabetLength]
	if x == nil || x.link != nil {
		return false
	}
	return x.key == s
}
