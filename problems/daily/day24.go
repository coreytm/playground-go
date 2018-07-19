package daily

// This problem was asked by Google.

// Implement locking in a binary tree.
// A binary tree node can be locked or unlocked only if all of its descendants or ancestors are not locked.

// Design a binary tree node class with the following methods:

// is_locked, which returns whether the node is locked
// lock, which attempts to lock the node. If it cannot be locked, then it should return false. Otherwise, it should lock it and return true.
// unlock, which unlocks the node. If it cannot be unlocked, then it should return false. Otherwise, it should unlock it and return true.

// You may augment the node to add parent pointers or any other property you would like.
// You may assume the class is used in a single-threaded program, so there is no need for actual locks or mutexes.
// Each method should run in O(h), where h is the height of the tree.

type lockableNode struct {
	parent              *lockableNode
	left                *lockableNode
	right               *lockableNode
	isLocked            bool
	numLockedDescendant int
}

func isLocked(node *lockableNode) bool {
	return node.isLocked
}

/*
	Simple solution to this problem is to keep track of whether a node is locked, and whether or not it has any locked descendants.
	By keeping track of locked descendants, when performing a lock or unlocke operation it won't have to iterate over all children
	which could be O(n) runtime.  Thus resulting on max O(2h) => O(h) runtime, where h is the height of the tree.

	To lock or unlock a node, all that needs to be done is to check if the node has any locked descendants,
	and to bubble up and check if all parents are unlocked.

	When a node is locked it notifies all of it's parents to increase it's count of locked descendants.
	When a node is unlocked it notifies all of it's parents to decrease it's count of locked descendants.
*/

// lock will lock the node and return true if all of the node's descendants and ancestors are unlocked.
// If any of the node's descendants or ancestors are locked it will return false.
//
// runtime: O(2h) where h is height of tree.  2h because if the node is locked, it iterates over parents again to notify of locked descendant.
// space: O(1)
func lock(node *lockableNode) bool {
	if node == nil || node.numLockedDescendant > 0 {
		return false
	}

	cur := node.parent
	for cur != nil {
		if cur.isLocked {
			return false
		}
	}

	node.isLocked = true

	cur = node
	for cur != nil {
		cur.numLockedDescendant++
	}

	return true
}

// unlock will unlock the node and return true if all of the node's descendants and ancestors are unlocked.
// If any of the node's descendants or ancestors are locked it will return false.
//
// runtime: O(2h) where h is height of tree.  2h because if the node is unlocked, it iterates over parents again to notify of unlocked descendant.
// space: O(1)
func unlock(node *lockableNode) bool {
	if node == nil || node.numLockedDescendant > 0 {
		return false
	}

	cur := node.parent
	for cur != nil {
		if cur.isLocked {
			return false
		}
	}

	node.isLocked = false

	cur = node.parent
	for cur != nil {
		cur.numLockedDescendant--
	}

	return true
}
