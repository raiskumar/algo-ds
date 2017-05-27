package tree

import "fmt"

//http://www.techiedelight.com/print-bottom-view-of-binary-tree/

// Generate Bottom view of the Binary tree
// Uses map to store map of relative position
// root is mapped to value 0; and then left node gets -1 and right node gets +1
// For every node, it's left and right children are given relative position
func bottomView(root *Node) {
	if root == nil {
		return
	}

	m := make(map[int]int)
	markNodes(root, m, 0)

	for k, v := range m {
		fmt.Printf("\n key[%d] value[%d]\n", k, v)
	}
}

// Mark nodes
func markNodes(root *Node, m map[int]int, mark int) {
	if root == nil {
		return
	}
	m[mark] = root.val // Only most recent values will be retained in the map
	markNodes(root.left, m, mark-1)
	markNodes(root.right, m, mark+1)
}
