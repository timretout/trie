package trie

import "testing"

func TestEmptyExists(t *testing.T) {
	tr := New()

	if tr.Exists("CAT") {
		t.Errorf("CAT should not be in trie")
	}
}

func TestExists(t *testing.T) {
	tr := New()

	tr.Insert("CAT")

	if !tr.Exists("CAT") {
		t.Errorf("CAT should be in trie")
	}
}

func TestNonExists(t *testing.T) {
	tr := New()

	tr.Insert("CAT")

	if tr.Exists("DOG") {
		t.Errorf("DOG should not be in trie")
	}
}

func TestSubstringNonExists(t *testing.T) {
	tr := New()

	tr.Insert("CATS")

	if tr.Exists("CAT") {
		t.Errorf("CAT should not be in trie")
	}
}

func TestInsert(t *testing.T) {
	tr := New()

	tr.Insert("CAT")
	tr.Insert("CATS")

	if !tr.Exists("CATS") {
		t.Errorf("CATS should be in trie")
	}
}

func TestDupeInsert(t *testing.T) {
	tr := New()

	tr.Insert("CAT")
	tr.Insert("CAT")

	if !tr.Exists("CAT") {
		t.Errorf("CAT should be in trie")
	}
}

func TestInsertSubstring(t *testing.T) {
	tr := New()

	data := []string{
		"CATASTROPHES",
		"CATASTROPHE",
		"CATS",
		"CAT",
	}

	for _, v := range data {
		tr.Insert(v)
	}

	for _, v := range data {
		if !tr.Exists(v) {
			t.Errorf("%s should be in trie", v)
		}
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
