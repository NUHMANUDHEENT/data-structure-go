package main

import (
	"fmt"
)

type TrieNode struct {
	Children map[rune]*TrieNode
	Isend    bool
}
type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			Children: map[rune]*TrieNode{}}}
}
func (t *Trie) Insert(word string) {
	node := t.root
	for _, v := range word {
		if _, exist := node.Children[v]; !exist {
			node.Children[v] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}
		node = node.Children[v]
	}
	node.Isend = true
}
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, v := range word {
		if _, exist := node.Children[v]; !exist {
			return false
		}
		node = node.Children[v]
	}
	return node.Isend
}
func (t *Trie) Prefix(word string) bool {
	currentNode := t.root
	for _, v := range word {
		if _, exist := currentNode.Children[v]; !exist {
			return false
		}
		currentNode = currentNode.Children[v]
	}
	return true
}
func (t *Trie) Delete(word string) {
	t.deleteHelper(t.root, word, 0)
}

func (t *Trie) deleteHelper(currentNode *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !currentNode.Isend {
			return false
		}
		currentNode.Isend = false
		// Return true if the current node has no other children (can be deleted)
		return len(currentNode.Children) == 0
	}

	char := rune(word[index])
	node, exists := currentNode.Children[char]
	if !exists {
		return false
	}

	// Recursively delete the child nodes
	shouldDelete := t.deleteHelper(node, word, index+1)

	if shouldDelete {
		delete(currentNode.Children, char)
		// Return true if the current node has no other children and is not the end of another word
		return len(currentNode.Children) == 0 && !currentNode.Isend
	}
	return false
}

func main() {

	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	fmt.Println("Search for 'hello':", trie.Search("helo"))
	trie.Delete("hello")
	fmt.Println("Search for 'hell':", trie.Prefix("hell"))
	// fmt.Println("StartsWith 'hell':", trie.StartsWith("hell")) // true
	// fmt.Println("StartsWith 'wor':", trie.StartsWith("wor"))   // true
	// fmt.Println("StartsWith 'word':", trie.StartsWith("word")) // false
}

