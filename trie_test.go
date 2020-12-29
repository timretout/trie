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

func TestDupeInsert(t *testing.T) {
	tr := New()

	tr.Insert("cat")
	tr.Insert("cat")

	if !tr.Exists("cat") {
		t.Errorf("cat should be in trie")
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
