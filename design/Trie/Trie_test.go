package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}