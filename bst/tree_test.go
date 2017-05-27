//How to Test : http://blog.alexellis.io/golang-writing-unit-tests/
// Go to the package directory and run below command
// $go test
package bst

import (
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
