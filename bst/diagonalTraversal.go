// Perform diagonal Traversal of the tree
//               1
//           /      \
//         2          3
//       /   \         \
//		4	  5			6
//
//  Diagonal View = 1,3,6
//                  2,5
//                   4

package bst

import "fmt"

func diagonalTraversal(root *Node) {
	if root == nil {
		return
	}

	//Map which will map diagonal number to corresponding list of values
	result := make(map[int][]int)

	diagonal(root, result, 0) // 3rd argument is to number diagonal

	// Print the values
	for k, v := range result {
		fmt.Printf("\n key[%d] value[%v]", k, v)
	}
}

//Key in the map represents the diagonal of the tree and its value stores all the diagonal elements
// DO a preorder traversal of the tree and populate the hash
// For each node, recursion is performed for left subtree by incrementing diagonal by 1. For right subtree the diagonal number remains same
func diagonal(root *Node, result map[int][]int, diagonalNumber int) {
	if root == nil {
		return
	}
	result[diagonalNumber] = append(result[diagonalNumber], root.val) //append current value to the existing value
	diagonal(root.left, result, diagonalNumber+1)
	diagonal(root.right, result, diagonalNumber)
}
