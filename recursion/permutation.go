package recursion

import "fmt"

// Given a String containing all DISTINCT characters, find all permuations of it
func permutation(input string) {
	if len(input) < 2 {
		return
	}
	output := "" // the permutated result gets stored
	generatePermutation(output, input)
}

// This approach will not work if there are duplicates in the String
// Time Complexity : O(N.N!)
//   Each permutation takes O(N); N is number of character in input string
//   And there are N! permutations
func generatePermutation(output, input string) {
	if len(input) == 0 {
		fmt.Println("perm: ", output)
		return
	}

	for i := 0; i < len(input); i++ {
		generatePermutation(output+input[i:i+1], input[0:i]+input[i+1:]) // won't work if this modification is done before metod call
	}
}
