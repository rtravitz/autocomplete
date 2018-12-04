package trie

import "strings"

//Build populates a trie with provided strings
func (t *Trie) Build(words []string) {
	populate(words, t.root)
}

func populate(words []string, n *node) {
	for _, word := range words {
		letters := strings.Split(word, "")
		if len(letters) > 0 {
			insertWord(letters, 0, n)
		}
	}
}

func insertWord(letters []string, idx int, n *node) {
	letter := letters[idx]
	if _, ok := n.children[letter]; !ok {
		n.children[letter] = createNode()
	}

	child := n.children[letter]

	if len(letters) == idx+1 {
		child.word = true
		return
	}

	insertWord(letters, idx+1, n.children[letter])
}
