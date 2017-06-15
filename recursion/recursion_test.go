package recursion

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	input := "ABC"
	permutation(input)
}

func TestAllBinaryStrings(t *testing.T) {
	input := "1?11?00?1?"
	allBinaryStrings(input)
}

func TestSubsetSum(t *testing.T) {
	a := []int{1, 2, 9, 3}
	subsetSum(a, 12)
}

func TestPowerSet(t *testing.T) {
	input := []int{1, 2, 3}
	generateAllSubsets(input)
}

func TestGenerateAllBinaryStrings(t *testing.T) {
	generateAllBinaryStrings(2)
}
