package recursion

// Calculate pow(x,n) i.e. x^n
// x^n = x*x*x.....*x, n of them

// This is NOT an efficient approach
// Uses recursion but not re-using already calcualted value
// Complexity : O(n)
func powerSlower(x, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}

	if n%2 == 0 {
		return power(x, n/2) * power(x, n/2)
	}
	return power(x, n/2) * power(x, n/2) * x
}

// This approach uses already calcualted value and size of the probelm reduces to half in each iteration
// Complexity: O(logn)
func power(x, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}

	pow := power(x, n/2)

	if n%2 == 0 {
		return pow * pow
	}
	return pow * pow * x
}
