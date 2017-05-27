package bst

// Calcualtes maximum sum from root node to any leaf node in Binary Tree or BST
// Time Complexity : O(N) or takes liner time
func maxSumRootToLeaf(root *Node, sum int, max *int) {
	if root == nil {
		if sum > *max { // If sum is greater than max so far then update max
			*max = sum
		}
		return
	}

	maxSumRootToLeaf(root.left, sum+root.val, max)
	maxSumRootToLeaf(root.right, sum+root.val, max)
}
