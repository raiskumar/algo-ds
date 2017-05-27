// Binary Search Tree Implementation
// Also provides implementation of common methods like height, traversal, min, max etc
//
//               100
//             /    \
//           50     200

package tree

import "fmt"

// Holds the reference to root node of BST
var root *Node

// Inserts a new value in BST at appropriate location
// After insertion, BST condition should holds good for all nodes (i.e. left node < root && right node > root)
func insert(v int, root *Node) *Node {
	if root == nil {
		return &Node{v, nil, nil}
	}

	if root.val > v {
		root.left = insert(v, root.left)
	} else {
		root.right = insert(v, root.right)
	}

	return root
}

// Performs inorder traversal of BST
// Inorder traversal prints the contents of BST in sorted order
func inorderTraversal(root *Node) {
	if root == nil {
		return
	}
	inorderTraversal(root.left)
	fmt.Println(root.val)
	inorderTraversal(root.right)
}

// Performs preorder traversal of BST
func preorderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.val)
	preorderTraversal(root.left)
	preorderTraversal(root.right)
}

// Performs PostOrder traversal of BST
func postOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	postOrderTraversal(node.left)
	postOrderTraversal(node.right)
	fmt.Println(node.val)
}

// Calculates height of the given BST
// Height = 1 + Max (Height of left subtree, Height of right subtree)
func height(node *Node) int {
	if node == nil {
		return 0
	}

	hl := height(node.left)  // calculate height of left subtree
	hr := height(node.right) // calculate height of right subtree

	if hl >= hr {
		return 1 + hl
	}
	return 1 + hr

}

// Calculates total number of nodes in given BST
// Total number of nodes = 1 + number of nodes in left subtree + number of nodes in right subtree
func getCountOfNodes(root *Node) int {
	if root == nil {
		return 0
	}

	return getCountOfNodes(root.left) + getCountOfNodes(root.right) + 1
}

// Takes array of elements and convert them to BST
// Returns root node of the BST
func createTree(elements []int) *Node {
	var root *Node
	for i, v := range elements {
		if i == 0 {
			root = insert(v, nil) //Creates root node, using the first element of the array
			continue
		}
		insert(v, root)
	}
	return root
}

// Find the node with maximum value
// It will be either the current node or the right most node in tree
// Complexity : O(h) Average case; h = height of tree
//              O(N) Worst case;  N = number of nodes
func findMaxNode(root *Node) *Node {
	if root == nil || root.right == nil {
		return root
	}
	return findMaxNode(root.right)
}

// Find the node with minimum value
func getMinNode(root *Node) *Node {
	if root == nil || root.left == nil {
		return root
	}
	return getMinNode(root.left)
}

// Get elements of tree in sorted order
func getSortedArray(root *Node) []int {
	var sortedVal []int
	inorder(root, sortedVal)
	return sortedVal
}

func inorder(root *Node, sortedVal []int) []int {
	if root == nil {
		return sortedVal
	}

	inorder(root.left, sortedVal)
	sortedVal = append(sortedVal, root.val)
	inorder(root.right, sortedVal)

	return sortedVal
}

// Find the node with the given value
func findNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if root.val == value {
		return root
	}
	if root.val > value {
		return findNode(root.left, value)
	} else if root.val < value {
		return findNode(root.right, value)
	}
	return nil
}

//Check if the Node is leaf node; i.e. does't have either of the children
func isLeaf(node *Node) bool {
	if node.left == nil && node.right == nil {
		return true
	}
	return false
}
