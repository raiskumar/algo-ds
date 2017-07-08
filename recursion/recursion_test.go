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

func TestKLengthSubset(t *testing.T) {
	kLengthSubset([]int{1, 2, 3, 4}, 3)
}

func TestCombinationSum(t *testing.T) {
	combinationSum([]int{2, 3, 6, 7}, 7)
}

func TestRatInMaze(t *testing.T) {
	maze := [][]int{
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{0, 1, 1, 1},
		{0, 0, 1, 1},
	}
	solution := ratInMaze(maze)
	fmt.Print("Solution Maze= %d", solution)
}

func TestGrayCode(t *testing.T) {
	solution := grayCode(3)
	fmt.Print("Solution = ", solution)
}
