//BST Problem
//If given key exists in BST then both floor and ceil is equal to that key
//Else Floor is previous greater key (if any) in BST and Ceil is next greater key (if any)

package tree

func getFloorAndCeil(root *Node, val int) (int, int) {
	if root == nil {
		return -1, -1
	}

	node := findNode(root, val)
	if node != nil { // if node is found in the BST
		return node.val, node.val
	}

	var floor, ceil int
	floorAndCeil(root, val, &floor, &ceil)
	return floor, ceil
}

//update current node to ceil if visiting left subtree and current node to floor if visiting right subtree
func floorAndCeil(root *Node, val int, floor, ceil *int) {
	if root == nil {
		return
	}

	if root.val < val {
		*floor = root.val
		floorAndCeil(root.right, val, floor, ceil)
	} else if root.val > val {
		*ceil = root.val
		floorAndCeil(root.left, val, floor, ceil)
	}
}
