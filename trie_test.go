package trie

import "testing"

func TestEmptyExists(t *testing.T) {
	tr := New()

	if tr.Exists("cat") {
		t.Errorf("cat should not be in trie")
	}
}

func TestExists(t *testing.T) {
	tr := New()

	tr.Insert("cat")

	if !tr.Exists("cat") {
		t.Errorf("cat should be in trie")
	}
}

func TestNonExists(t *testing.T) {
	tr := New()

	tr.Insert("cat")

	if tr.Exists("dog") {
		t.Errorf("dog should not be in trie")
	}
}

func TestSubstringNonExists(t *testing.T) {
	tr := New()

	tr.Insert("cats")

	if tr.Exists("cat") {
		t.Errorf("cat should not be in trie")
	}
}

func TestInsert(t *testing.T) {
	tr := New()

	tr.Insert("cat")
	tr.Insert("cats")

	if !tr.Exists("cats") {
		t.Errorf("cats should be in trie")
	}

}

func TestInsertEmptyString(t *testing.T) {
	tr := New()

	if tr.Exists("") {
		t.Errorf("empty string should not be in trie")
	}

	tr.Insert("")

	if !tr.Exists("") {
		t.Errorf("empty string should be in trie")
	}

}

func TestInsertZeroRune(t *testing.T) {
	tr := New()

	if tr.Exists("\x00") {
		t.Errorf("zero rune should not be in trie")
	}

	tr.Insert("\x00")

	if !tr.Exists("\x00") {
		t.Errorf("zero rune should be in trie")
	}

}

// BenchmarkTrie-8     	  669060	      1609 ns/op	    1120 B/op	      35 allocs/op
// BenchmarkExists-8   	100000000	        10.7 ns/op	       0 B/op	       0 allocs/op

func BenchmarkTrie(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := New()

		for _, v := range []string{
			"cat", "cats", "dog", "dogs", "caterpillar", "catenary", "catastrophe",
		} {
			tr.Insert(v)
		}
		tr.Exists("cats")
	}
}

func BenchmarkExists(b *testing.B) {
	tr := New()

	for _, v := range []string{
		"cat", "cats", "dog", "dogs", "caterpillar", "catenary", "catastrophe",
	} {
		tr.Insert(v)
	}

	for n := 0; n < b.N; n++ {
		tr.Exists("cats")
	}
}
