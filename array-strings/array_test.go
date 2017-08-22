package array

import (
	"fmt"
	"testing"
)

func TestGetMaxProductsOfAllIntsExceptOne(t *testing.T) {
	var arr = []int{2, 4, 9, 3, -100}
	result := getMaxProductsOfAllIntsExceptOne(arr)
	fmt.Println("result =", result)
}
