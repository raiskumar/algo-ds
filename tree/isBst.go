//BST Property
//Every node on the right subtree has to be larger than the current node 
// and every node on the left subtree has to be smaller than the current node
package tree

//Assume that BST has elements in below two ranges
const maxInt = 32767 // max value of 16 bit integer
const minInt = -32768

// This approach is NOT very efficient
// Each node should be greater than the largest value found in left sub tree and smaller than smallest value found in right sub tree
// Complexity: findMaxNode has O(h)
//             Each node of the tree needs to call max as well as min method i.e. O(h) for each node
//             O(h*N)
func isBst(root *Node) bool {
	if root == nil {
		return true
	}

	if root.left != nil && findMaxNode(root.left).val > root.val { // findMax is implemented in bst.go
		return false
	}
	if root.right != nil && findMinNode(root.right).val < root.val {
		return false
	}

	if !isBst(root.left) || !isBst(root.right) {
		return false
	}

	return true
}

//Perform inorder traversal and store all elements in an array/list
//Scan once to check if array is sorted
//This approach has an overhead of extra space i.e. O(N)
func isBstUsingInOrderTraversal(root *Node) bool {
	arr := getSortedArray(root)

	var tmp int
	for i, v := range arr {
		if i == 0 {
			tmp = arr[i]
			continue
		}
		if tmp > v {
			return false
		}

		tmp = arr[i-1]
	}

	return true
}

// Optimized approach : traverses each node only once
// Keep track of minimum and maximum allowed value for each node and breaks the look if this condition is violated
// This approach also has space overhead of O(N); but it's not EXPLICIT
// Time complexity ; O(N)
func isBstImproved(root *Node) bool {
	if root == nil {
		return true
	}

	return isBstImproved2(root, minInt, maxInt) // oops go doesn't support method overloading!
}

func isBstImproved2(root *Node, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.val > minVal && root.val < maxVal {
		return isBstImproved2(root.left, minVal, root.val) && isBstImproved2(root.right, root.val, maxVal)
	}

	return false
}
