package tree

import "fmt"

func getInorderSuccessor(root *Node, target int) *Node {
	node := findNode(root, target)
	//return successor(root, node)
	result := node
	successorV2(root, node, result)
	return result
	//return successorV3(root, node)
}

func successor(root, target *Node) *Node {
	if target == nil || root == nil {
		return nil
	}

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
