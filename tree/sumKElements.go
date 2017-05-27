package tree

//Sum first K elements in inorder traversal of the BST / Binary TREE
func sumFirstKElements(root *Node, k *int) int {
	if root == nil {
		return 0
	}

	//sum := 0
	sum := sumFirstKElements(root.left, k) // compute sum of elements in left subtree
	if *k == 0 {
		return sum
	}
	*k = *k - 1
	sum = sum + root.val
	return sum + sumFirstKElements(root.right, k)
}

func sumLastKElements(root *Node, k *int) int {
	if root == nil {
		return 0
	}

	sum := sumLastKElements(root.right, k)
	if *k == 0 {
		return sum
	}

	*k = *k - 1
	sum = sum + root.val

	return sum + sumLastKElements(root.left, k)
}
