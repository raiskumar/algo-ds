package dp

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	n := 5
	fmt.Printf("fibonacci of %d = %d", n, fibonacci(n))
}

func TestMaxSumSubarray(t *testing.T) {
	arr := []int{904, 40, 523, 12, -335, -385, -124, 481, 31}
	i, j, sum := maxSubArray(arr)
	fmt.Printf("max Sum between %d,%d = %d", i, j, sum)
}
