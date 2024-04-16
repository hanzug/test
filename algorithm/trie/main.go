package main

import "fmt"

const (
	alphabetSize = 26
)

type Node struct {
	child       [alphabetSize]*Node
	isEndofWord bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(w string) {
	node := t.root
	for _, ch := range w {
		if node.child[ch-'a'] == nil {
			node.child[ch-'a'] = &Node{}
		}
		node = node.child[ch-'a']
	}
	node.isEndofWord = true
}

func (t *Trie) Search(w string) bool {
	node := t.root
	for _, ch := range w {
		if node.child[ch-'a'] == nil {
			return false
		}
		node = node.child[ch-'a']
	}
	if node.isEndofWord == true {
		return true
	} else {
		return false
	}
}

func main() {
	trie := NewTrie()
	words := []string{"the", "a", "there", "answer", "any", "by", "their"}

	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("the"))   // Output: true
	fmt.Println(trie.Search("these")) // Output: false
	fmt.Println(trie.Search("their")) // Output: true
	fmt.Println(trie.Search("ther"))  // Output: false
}
