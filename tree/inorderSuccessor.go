package tree

import "fmt"

//Approache1
// Do an inorder walk and getting all elements in the array. Then find the node in the array in one scan and return the next node. 
// Doesn't use BST Property

// Use BST search idiom

func successor(root, target *Node) *Node {
	if target == nil || root == nil {
		return nil
	}

	// If target node has right subtree
	if target.right != nil {
		return getMinNode(target.right)
	}

	var result *Node
	for root != nil {
		if target.val < root.val {
			result = root
			root = root.left
		} else if target.val > root.val {
			root = root.right
		} else {
			break
		}
	}
	return result
}

func getInorderSuccessor(root *Node, target int) *Node {
	node := findNode(root, target)
	result := node
	successorV2(root, node, result)
	return result
}


func successorV2(root, target, result *Node) {
	if target == nil || root == nil {
		return
	}

	if root.val < target.val {
		result = root
		fmt.Println(" result ", result)
		successorV2(root.left, target, result)
	} else {
		successorV2(root.right, target, result)
	}
	return
}

//Iterative version
func successorV3(root, target *Node) *Node {
	if target == nil || root == nil {
		return nil
	}

	var result *Node
	for root != nil {
		if root.val > target.val {
			result = root
			root = root.left
		} else {
			root = root.right
		}
	}
	return result
}
