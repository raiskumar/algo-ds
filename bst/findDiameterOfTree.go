package bst

//Reference : http://www.techiedelight.com/find-diameter-of-a-binary-tree/

func findDiameter(root *Node) int {
	if root == nil {
		return 0
	}

	d := 0
	//diameter(root, &d)
	println(" diameter of tree =", d)
	println("diameter ", diameterv2(root, &d))

	return d
}

//Calculates height of left subtree and right subtree for each node of the tree.
//Diameter for a node will be height of left and height of right subtree +1
//Final, diameter of the tree will be the maximum of all diameters
// Complexity is O(N^2)
func diameter(root *Node, d *int) {
	if root == nil {
		return
	}

	//Get height of left and right subtree
	hl := height(root.left)
	hr := height(root.right)

	//Update diameter if the current one is smaller than the new one
	if hl+hr+1 > *d {
		*d = hl + hr + 1
	}

	//calcualte diameter of left and right subtree
	diameter(root.left, d)
	diameter(root.right, d)
}

//In above approach; height is getting calculated separately
// Start at the bottom of the tree and then return height of the subtree rooted at given node to its parent
// This is optimized approach, will take linear time
func diameterv2(root *Node, diameter *int) int {
	if root == nil {
		return 0
	}

	hl := diameterv2(root.left, diameter)
	hr := diameterv2(root.right, diameter)

	max_diameter := hl + hr + 1

	if max_diameter > *diameter {
		diameter = &max_diameter
	}
	if hl > hr {
		return hl + 1
	}
	return hr + 1
}
