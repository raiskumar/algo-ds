package bst

//Find 2nd largest node of the BST
//Takes root of the tree as the input
func secondLargestNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.right == nil {
		return findMaxNode(root.left) // if there is no right subtree; then largest of left subtree will be 2nd largest.
	}
	if root.right != nil {
		if isLeaf(root.right) {
			return root
		}
	}
	return secondLargestNode(root.right)
}
