// Two nodes of BST are swapped, correct the BST

package tree

import "fmt"

func correctSwappedNodes(root *Node) bool {
	if root == nil {
		return true
	}
	min, max := -1, 10000
	var first, second *Node //&Node{-1, nil, nil}, &Node{-1, nil, nil}
	fixBst(root, first, second, min, max)
	inorderTraversal(root)
	return true
}

func fixBst(root, first, second *Node, min, max int) {
	if root == nil {
		return
	}
	if first != nil && second != nil {
		return
	}
	fmt.Println(first, " - ", second)
	if root.val < max && root.val > min {
		fixBst(root.left, first, second, min, root.val)
	} else {
		if first == nil {
			fmt.Println("first is null ", root.val, first)
			first = &Node{root.val, root.left, root.right}
			fmt.Println("first is null-- ", (first == nil), first)
			return
		} else {
			second = &Node{root.val, root.left, root.right} //&Node{root.val, root.left, root.right}
			first.val, second.val = second.val, first.val
			return
		}
	}

	if root.val < max && root.val > min {
		fixBst(root.right, first, second, root.val, max)
	} else {
		if first == nil {
			fmt.Println("first is null ", root.val)
			first = &Node{root.val, root.left, root.right}
			fmt.Println("first is null-- ", first)
			return
		} else {
			second = &Node{root.val, root.left, root.right} //&Node{root.val, root.left, root.right}
			first.val, second.val = second.val, first.val
			return
		}
	}
}
