package recursion

import "fmt"

// Problem Statement: Generate All possible Binary Strings of length k
// If k = 1; Result= 0 , 1
// If k = 2; Result= 00, 01, 10, 11
func generateAllBinaryStrings(k int) {
	if k == 0 {
		return
	}

	result := ""
	generateAllBinaries(k, 0, result)
}

// Start with empty result stering and for each index append 0 and 1
// And once length of the result is 3, we have got one binary string
func generateAllBinaries(k, index int, result string) {
	if len(result) == k { // Or index+1 == len
		fmt.Println(result)
		return // Can't Miss return
	}

	generateAllBinaries(k, index+1, result+"0")
	generateAllBinaries(k, index+1, result+"1")
}
