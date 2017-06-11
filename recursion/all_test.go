package recursion

import (
	"testing"
)

func TestKPartition(t *testing.T) {
	var elements = []int{100, 50, 200, 150, 140, 160, 170, 130, 120}
	kPartition(elements, 2)
}

func TestPermutation(t *testing.T) {
	input := "ABC"
	permutation(input)
}
