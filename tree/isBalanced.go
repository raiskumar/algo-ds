package tree

import "github.com/raiskumar/algo-ds/common"

//Approach 1  (Inefficient)
// Calculate the height of left subtree and right subtee for each node
// For any node if the difference in height of left and right subtree is > 1; tree is UNBALANCED

//Approach 2
// Do post order traversal and strt from bottom
// Return height of subtree rooted at given node to its parent node

// Approach 2
func isHeightBalanced(root *Node, isBalanced *bool) int {
	if root == nil {
		return 0
	}

	heightLeft := isHeightBalanced(root.left, isBalanced)
	heightRight := isHeightBalanced(root.right, isBalanced)

	if common.Abs(heightLeft-heightRight) > 1 {
		v := false
		isBalanced = &(v)
	}

	return common.Max(heightLeft, heightRight) + 1
}
