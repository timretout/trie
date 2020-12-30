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

	tr.Insert("cat", 1)

	if !tr.Exists("cat") {
		t.Errorf("cat should be in trie")
	}
}

func TestNonExists(t *testing.T) {
	tr := New()

	tr.Insert("cat", 1)

	if tr.Exists("dog") {
		t.Errorf("dog should not be in trie")
	}
}

func TestSubstringNonExists(t *testing.T) {
	tr := New()

	tr.Insert("cats", 1)

	if tr.Exists("cat") {
		t.Errorf("cat should not be in trie")
	}
}

func TestInsert(t *testing.T) {
	tr := New()

	tr.Insert("cat", 1)
	tr.Insert("cats", 1)

	if !tr.Exists("cats") {
		t.Errorf("cats should be in trie")
	}
}

func TestGet(t *testing.T) {
	tr := New()

	tr.Insert("cat", 1234)
	v := tr.Get("cat")
	if v != 1234 {
		t.Errorf("cat should be 1234, got %d", v)
	}

	tr.Insert("cat", 5678)
	v = tr.Get("cat")
	if v != 5678 {
		t.Errorf("cat should be 5678, got %d", v)
	}
}

func TestDupeInsert(t *testing.T) {
	tr := New()

	tr.Insert("cat", 1)
	tr.Insert("cat", 2)

	if !tr.Exists("cat") {
		t.Errorf("cat should be in trie")
	}
}

func TestInsertEmptyString(t *testing.T) {
	tr := New()

	if tr.Exists("") {
		t.Errorf("empty string should not be in trie")
	}

	tr.Insert("", 1)

	if !tr.Exists("") {
		t.Errorf("empty string should be in trie")
	}
}

func TestInsertZeroRune(t *testing.T) {
	tr := New()

	if tr.Exists("\x00") {
		t.Errorf("zero rune should not be in trie")
	}

	tr.Insert("\x00", 1)

	if !tr.Exists("\x00") {
		t.Errorf("zero rune should be in trie")
	}
}
