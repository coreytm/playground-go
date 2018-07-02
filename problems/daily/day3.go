package daily

import (
	"fmt"
	"strings"
)

// This problem was asked by Google.
// Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.
// For example, given the following Node class

// class Node:
//    def __init__(self, val, left=None, right=None):
//        self.val = val
//        self.left = left
//        self.right = right
// The following test should pass:

// node = Node('root', Node('left', Node('left.left')), Node('right'))
// assert deserialize(serialize(node)).left.left.val == 'left.left'

type node struct {
	val   string
	left  *node
	right *node
}

// serialize will convert the tree into a poorly formatted json string where new lines are not indented.
// This is a recursive function.
//
// runtime: O(n)
// space: O(1)
func serialize(root *node) string {
	if root == nil {
		return "null"
	}

	left := serialize(root.left)
	right := serialize(root.right)

	return fmt.Sprintf(`{\n"val":"%s",\n"left":%s,\n"right":%s}`, root.val, left, right)
}

// deserialize accepts the serialized string of a tree (output from serialize() function) and returns the root of a tree.
// This function starts by splitting the string up using a newline character as a delimiter and then
// calls the recursive deserializeRecurive() function to perform the deserialization recursively.
//
// runtime: O(n) for splitting the string
// space: O(n) for creating an []string
func deserialize(s string) *node {
	if s == "null" || s == "" {
		return nil
	}

	parts := strings.Split(s, "\n")
	root, _ := deserializeRecursive(parts, 0)
	return root
}

// deserializeRecursive will iterate over each string and look for val, left, and right.
// When it finds a '}' it will return the node, and the current "line" or string position it's at.
// This way the calling function can continue where the recurive calls left off.
//
// runtime: O(n)
// space: O(1)
func deserializeRecursive(in []string, curLine int) (*node, int) {
	n := len(in)
	if n == 0 {
		return nil, curLine
	}

	root := node{}
	for curLine < n && in[curLine] != "}" {
		if strings.Contains(in[curLine], `"val":`) {
			root.val = strings.Split(in[curLine], `"`)[2]
		}
		if strings.Contains(in[curLine], `"left":{`) {
			root.left, curLine = deserializeRecursive(in, curLine+1)
		}
		if strings.Contains(in[curLine], `"right":{`) {
			root.right, curLine = deserializeRecursive(in, curLine+1)
		}

		curLine++
	}

	return &root, curLine + 1
}
