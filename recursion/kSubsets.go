package recursion

import "fmt"

// Generate k length subsets of set {1,2,...}
// Mathematically expressed as nCk
// Brute force solution : Generate power set as implemented in  recursion/powerSet.go and then filter out subsets of length k

// Pascal Identity : (n,k) = (n-1,k)+(n-1, k-1)
//                     nCk = n-1Ck + n-1Ck-1

func kLengthSubset(set []int, k int) {
	if k == 0 {
		return
	}

	var result []int
	generateKLengthSubset(set, result, k)
}

// For every element of input we have two options, whether to select it or ignore it
// Perform this recursively until k is 0
func generateKLengthSubset(input, output []int, k int) {
	if k == 0 {
		fmt.Printf("\n subset =%v", output)
		return
	}

	if len(input) > 0 {
		val := input[0]
		if k > 0 {
			//The current element is chosen . Now we recursively choose the remaining k-1 elements from the remaining set
			generateKLengthSubset(input[1:], append(output, val), k-1)

			//The current element is not chosen . Now we recursively choose k elements from the remaining set.(exclusion)
			generateKLengthSubset(input[1:], output, k)
		}
	}
}
