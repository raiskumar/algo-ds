package bst

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	var elements = []int{100, 50, 200, 150, 140, 160, 170, 130, 120}
	root := createTree(elements)
	fmt.Println(" root =", root)
	sum := 0
	max := 0
	maxSumRootToLeaf(root, sum, &max)
	println("sum =", max)
	if max != 840 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 840)
	}
}
