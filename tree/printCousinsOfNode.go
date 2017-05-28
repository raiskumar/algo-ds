package tree

// We can't use BST property to solve this problem;
// It's a Binary Tree problem
func printAllCousins(root, node *Node) {
	level, index := 1, 1
	getLevelOfNode(root, node, index, &level)
	println("level of node = ", level)
	printCousins(root, node, level)
}

//Find level of given node by doing pre-order traversal of the tree.
func getLevelOfNode(root, node *Node, index int, level *int) {
	if root == nil || node == nil {
		return
	}

	if root == node {
		*level = index
		println("level =", *level)
		return
	}

	getLevelOfNode(root.left, node, index+1, level)
	getLevelOfNode(root.right, node, index+1, level)
}

// Print all nodes at the level excluding node and it's sibling
func printCousins(root, node *Node, level int) {
	if root == nil {
		return
	}

	if level == 1 { //print node value if level is reduced to 1.
		println(root.val)
	}

	//Recurse down left and right if root is not parent of the node
	if root.left != node && root.right != node {
		printCousins(root.left, node, level-1)
		printCousins(root.right, node, level-1)
	}

}
