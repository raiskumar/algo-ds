package dp

// Generate fibonacci sequence for n
func fibonacci(n int) int {
	//return fibonacciBottomUp(n)
	return fibonacciMemoization(n)
}

// Stores last 2 values and then uses it to arrive at the current val
func fibonacciBottomUp(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	secondLastVal := 0
	lastVal := 1
	val := 0

	for i := 2; i < n; i++ {
		val = secondLastVal + lastVal
		secondLastVal = lastVal
		lastVal = val
	}
	return val
}

// Stores all values in array upto n
func fibonacciMemoization(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	result := make([]int, n+1)
	result[0] = 0
	result[1] = 1

	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n-1]
}
