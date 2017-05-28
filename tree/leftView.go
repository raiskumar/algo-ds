package tree

import "fmt"

//http://www.techiedelight.com/print-left-view-of-binary-tree/

// Prints Left view of the binary tree
// It prints only the first left most element from each level
// Not an efficient solution as this method gets called for each level
func leftView(root *Node) {
	if root == nil {
		return
	}

	height := height(root) // total number of levels
	processed := false

	for i := 1; i <= height; i++ {
		printLeveViewAtGivenLevel(root, &processed, i)
		processed = false // re-initialize the flag to false
	}
}

// Print the left most node at given level
// processed controls the loop; the recursion stops once the value become false
func printLeveViewAtGivenLevel(root *Node, processed *bool, level int) {
	if root == nil || *processed == true {
		return
	}
	if level == 1 {
		*processed = true
		println(" left view ", root.val)
	}

	printLeveViewAtGivenLevel(root.left, processed, level-1)
	printLeveViewAtGivenLevel(root.right, processed, level-1)
}

// 2nd more efficient version
// Uses map to keep track if left most node is already considered or not
func leftViewImproved(root *Node) {
	if root == nil {
		return
	}

	//var m map[int]int
	m := make(map[int]int)

	lvImproved(root, m, 1)

	for level, value := range m {
		fmt.Printf(" \n level =%d & value =%d", level, value)
	}
}

func lvImproved(root *Node, m map[int]int, level int) {
	if root == nil {
		return
	}

	_, ok := m[level] // return value if found or ok=false if not found

	if ok == false {
		m[level] = root.val
	}

	lvImproved(root.left, m, level+1)
	lvImproved(root.right, m, level+1)
}

func leftViewImproved2(root *Node) {
	if root == nil {
		return
	}

	currLevel, lastLevel := 1, 0
	lvI2(root, currLevel, &lastLevel)
}

func lvI2(root *Node, cur int, last *int) {
	if root == nil {
		return
	}

	if *last < cur {
		fmt.Printf(" \n level =%d & value =%d", cur, root.val)
		*last = cur
	}

	lvI2(root.left, cur+1, last)
	lvI2(root.right, cur+1, last)
}
