package array

import (
	"fmt"
	"testing"
)

func TestGetMaxProductsOfAllIntsExceptOne(t *testing.T) {
	var arr = []int{2, 4, 9, 3, -100, 1}
	result := getMaxProductsOfAllIntsExceptOne(arr)
	fmt.Println("result =", result)

	result1 := getMaxProductsOfAllIntsExceptOneOptimized(arr)
	fmt.Println("result =", result1)
}

func TestLongestContiguousIncreasingSubarray(t *testing.T) {
	var arr = []int{2, 11, 3, 5, 13, 7, 19, 17, 233}
	index, size := longestContiguousIncreasingSubarray(arr)
	fmt.Println("(i,j)=", index, index+size)

	var arr2 = []int{10, 5, 3, 10}
	index, size = longestContiguousIncreasingSubarray(arr2)
	fmt.Println("(i,j)=", index, index+size)
}

func TestRotate(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(input, 3)
	fmt.Println("%v", input)

	input2 := []int{1, 2, 3, 4, 5, 6, 7}
	rotateOptimzed(input2, 3)
	fmt.Println("%v", input2)
}
