package dp

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	n := 5
	fmt.Printf("fibonacci of %d = %d", n, fibonacci(n))
}
