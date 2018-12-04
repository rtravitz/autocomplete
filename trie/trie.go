package trie

//Trie is the root of a search tree
type Trie struct {
	root *node
}

//New returns a pointer to a Trie
func New() *Trie {
	return &Trie{root: createNode()}
}

type node struct {
	word     bool
	children map[string]*node
}

func createNode() *node {
	children := make(map[string]*node)
	return &node{word: false, children: children}
}
