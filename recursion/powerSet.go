package recursion

import "fmt"

// Generate Power Set of a set, S
// Power Set is all subsets of set including both the empty set and S itself
// S = {1,2,3}
// Power Set = {}, {1}, {2}, {3}, {1,2}, {2,3}, {1,3}, {1,2,3}
// Size = pow(2,N); here N = 3
func generateAllSubsets(input []int) {
	var result []int //Empty slice; it can grow as well
	subsets(input, result)
}

// For each element of the input there are two possibilites :
//  1. Either it's NOT part of the result
//  2. or, It's part
func subsets(input, result []int) {
	if len(input) == 0 {
		fmt.Printf("\n Subset =%v", result)
		return
	}

	val := input[0]
	subsets(input[1:], result)              //val not part of result
	subsets(input[1:], append(result, val)) // val part of result
}
