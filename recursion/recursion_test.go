package recursion

import (
	"fmt"
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

func TestPatterMatch(t *testing.T) {
	//patternMatch("a*b.c*abcd", "abc")
	fmt.Println(patternMatch("a*b.c", "abcc"))
	fmt.Println(patternMatch("a*b.", "abc"))
}

func TestPatterMatchV2(t *testing.T) {
	fmt.Println(patternMatchV2("a*b.c", "aaaabdc"))
	fmt.Println(patternMatchV2("a*b.c", "bcc"))
}

func TestPower(t *testing.T) {
	fmt.Println(powerSlower(2, 5))
	fmt.Println(power(2, 3))
}
