package algorithm

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	tree := new(TrieTree)
	tree.Add("TypeScript")
	tree.Add("Python")
	tree.Add("Go言語")
	tree.Add("Golang")

	terms := tree.Find("Go")
	fmt.Println(terms)
	terms = tree.Find("T")
	fmt.Println(terms)

}
