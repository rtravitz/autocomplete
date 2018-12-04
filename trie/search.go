package trie

import "strings"

//Suggest returns a slice of suggestions for a given substring
func (t *Trie) Suggest(substring string) []string {
	letters := strings.Split(substring, "")
	startingNode := reachStartingNode(letters, 0, t.root)
	suggestions := searchChildren(startingNode, "", []string{}, substring)
	return suggestions
}

func searchChildren(n *node, suffix string, suggestions []string, prefix string) []string {
	if n.word == true {
		word := prefix + suffix
		suggestions = append(suggestions, word)
	}

	if !(len(n.children) == 0) {
		for letter, child := range n.children {
			newSuffix := suffix + letter
			suggestions = searchChildren(child, newSuffix, suggestions, prefix)
		}
	}

	return suggestions
}

func reachStartingNode(letters []string, idx int, n *node) *node {
	if len(letters) == idx {
		return n
	}

	letter := letters[idx]
	if child, ok := n.children[letter]; ok {
		return reachStartingNode(letters, idx+1, child)
	}

	return nil
}
