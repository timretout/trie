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
