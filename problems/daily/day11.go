package daily

import (
	"math"
	"strings"
)

// This problem was asked by Twitter.
// Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.
// For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].
// Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.

// autoComplete will return all the strings in 'in' that start with 's'.
// this function does it using the iterative approach and the strings library.
//
// runtime: O(n) because it iterates over all strings, however performance of strings.HasPrefix() might affect this.
// space:   O(n) because of out slice
func autoComplete(s string, in []string) []string {
	out := []string{}
	for _, str := range in {
		if strings.HasPrefix(s, str) {
			out = append(out, str)
		}
	}

	return out
}

type trieNode struct {
	val      byte
	children map[byte]*trieNode
	isEnd    bool
	word     string
}

func autoCompleteWithTrie(s string, words []string) []string {
	root := trieNode{
		children: map[byte]*trieNode{},
	}

	minLen := math.MaxInt32
	for _, word := range words {
		if len(word) < minLen {
			minLen = len(word)
		}

		insert(&root, word)
	}

	// if the substring is longer than all the provided words, then return empty slice
	if len(s) > minLen {
		return []string{}
	}

	// go to node of last letter in s
	cur := &root
	curLetter := 0
	for cur.children[s[curLetter]] != nil {
		cur = cur.children[s[curLetter]]
		curLetter++
	}

	// all paths from here with nodes where isEnd=true is a word which contains s
	queue := []*trieNode{}
	for _, v := range cur.children {
		queue = append(queue, v)
	}

	// perform a BFS to get the words
	out := []string{}
	for len(queue) > 0 {
		// dequeue
		top := queue[0]
		queue = queue[1:]

		for _, v := range top.children {
			queue = append(queue, v)
		}

		if top.isEnd {
			out = append(out, top.word)
		}
	}

	return out
}

// insert will insert the word into a trie.
func insert(root *trieNode, word string) {
	tn := root.children[word[0]]
	if tn == nil {
		tn = &trieNode{
			val:      word[0],
			children: map[byte]*trieNode{},
			word:     string(word[0]),
		}

		root.children[word[0]] = tn
	}

	// save the word so it's easier to retrieve during BFS
	tn.word = root.word + string(word[0])

	if len(word) == 1 {
		root.children[word[0]].isEnd = true
		return
	}

	insert(root.children[word[0]], word[:len(word)])
}
