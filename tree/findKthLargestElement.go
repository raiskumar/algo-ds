package tree

// Brute-force Approach
// Perform inorder traversal which will return the content of tree in ascending manner
// Return the element from the kth position from last
// Not efficient - processes many nodes which are not required

// Reverse-Inorder traversal
// Recurse first in right subtree and then on left subtree and return as soon as k becomes 0
//Time Complexity : O(h+k)

//Reverse-Inorder Implementation
func findKthLargestElement(root *Node, k *int, response *int) {
	if root == nil || *k == 0 {
		return
	}

	findKthLargestElement(root.right, k, response)
	*k = *k - 1
	if *k == 0 {
		*response = root.val
		return
	}
	findKthLargestElement(root.left, k, response)
}
