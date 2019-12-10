package main

import (
	"fmt"

	"github.com/po3rin/trie"
)

func main() {
	words := []string{"go", "golang", "godoc", "google", "goto", "ago", "ゴー", "ゴーラン"}
	root := trie.GetNode()

	for i := 0; i < len(words); i++ {
		root.Insert(words[i])
	}

	fmt.Println("[google]: ", trie.Search(root, "google"))
	fmt.Println("[go]: ", trie.Search(root, "go"))
	fmt.Println("[lang]: ", trie.Search(root, "lang"))
	fmt.Println("[ゴーラング]: ", trie.Search(root, "ゴーラング"))
}
