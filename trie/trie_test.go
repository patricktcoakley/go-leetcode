package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Fatal("didn't find 'app'")
	}

	if trie.Search("app") {
		t.Fatal("found 'apple'")
	}

	if !trie.StartsWith("app") {
		t.Fatal("didn't find prefix 'app'")
	}

	trie.Insert("app")

	if !trie.Search("app") {
		t.Fatal("didn't find 'app'")
	}
}
