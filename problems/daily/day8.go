package daily

// This problem was asked by Google.
// A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.
// Given the root to a binary tree, count the number of unival subtrees.
// For example, the following tree has 5 unival subtrees:

//    0
//   / \
//  1   0
//     / \
//    1   0
//   / \
//  1   1

// numUnivalSubtrees returns the number of unival subtrees in the provided tree.
func numUnivalSubtrees(root *node) int {
	cnt, _ := numUnivalSubtreesRecursive(root)
	return cnt
}

// numUnivalSubtreesRecursive returns the current number of unival sub trees, and whether or not the current tree is a unival subtree.
// runtime: O(n) // iterate through each node, but adds to the call stack (height of tree)
// spaec:   O(1)
func numUnivalSubtreesRecursive(root *node) (int, bool) {
	cnt := 0
	if root == nil {
		return cnt, true
	}

	left, leftIsUST := numUnivalSubtreesRecursive(root.left)
	right, rightIsUST := numUnivalSubtreesRecursive(root.right)

	if leftIsUST && rightIsUST {
		cnt = left + right

		if root.left != nil && root.left.val != root.val {
			return cnt, false
		}

		if root.right != nil && root.right.val != root.val {
			return cnt, false
		}

		return cnt + 1, true
	}

	return cnt, false
}
