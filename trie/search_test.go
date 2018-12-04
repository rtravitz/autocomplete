package trie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/rtravitz/autocomplete/trie"
)

var _ = Describe("Search", func() {
	It("finds all words under the prefix", func() {
		trie := New()
		words := []string{"dog", "dogged", "cat"}
		trie.Build(words)
		result := trie.Suggest("do")
		expected := []string{"dog", "dogged"}
		Expect(result).To(Equal(expected))
	})
})
