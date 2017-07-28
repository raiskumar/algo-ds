package tree

import (
	"fmt"
	"testing"
)

func TestMaxSumToLeaf(t *testing.T) {
	var elements = []int{100, 50, 200, 150, 140, 160, 170, 130, 120}
	root := createTree(elements)
	sum := 0
	max := 0
	maxSumRootToLeaf(root, sum, &max)
	println("sum =", max)
	if max != 840 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 840) //t.Error or t.Fail to indicate failure
	}
}

func TestIsBalanced(t *testing.T) {
	var elements = []int{100, 50, 200, 150, 140, 160, 170, 130, 120}
	root := createTree(elements)
	balanced := false
	isHeightBalanced(root, &balanced)
	println("Tree is Balanced ? ", balanced)
}

func TestDiagonalTraversal(t *testing.T) {
	var elements = []int{100, 50, 200, 300, 60, 30}
	root := createTree(elements)
	diagonalTraversal(root)
}

func TestLevelOrderTraversal(t *testing.T) {
	var elements = []int{100, 50, 200, 150, 140, 160, 170, 130, 120}
	root := createTree(elements)
	levelOrderTraversal(root)
}

func TestKthLargestElement(t *testing.T) {
	var elements = []int{100, 50, 200, 150, 140, 160, 170, 130, 120}
	root := createTree(elements)

	response := -1
	k := 5
	findKthLargestElement(root, &k, &response)
	fmt.Printf(" kth largest = %d ", response)
}

func TestFloorAndCeil(t *testing.T) {
	var elements = []int{100, 50, 200, 150, 140, 160, 170, 130, 120}
	root := createTree(elements)

	floor, ceil := getFloorAndCeil(root, 110)
	fmt.Println(floor, " <- floor ceil -> ", ceil)
}

func TestCorrectTwoSwappedNodesOfBst(t *testing.T) {
	var elements = []int{100, 50, 200}
	root := createTree(elements)

	first := findNode(root, 50)
	second := findNode(root, 200)

	first.val, second.val = second.val, first.val
	inorderTraversal(root)

	fmt.Println("now nodes are corrupted")
	var one, two *Node
	fmt.Println(one)
	fmt.Println(two)
	one = &Node{5, nil, nil}
	two = &Node{5, nil, nil}
	fmt.Println("one ", one)
	fmt.Println(two == nil)

	correctSwappedNodes(root)
}
