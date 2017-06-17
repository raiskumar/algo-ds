package recursion

import "fmt"

// Find all combinations which will result in sum
func combinationSum(input []int, sum int) {
	if sum <= 0 {
		return
	}
	var result []int
	generateCombinations(input, result, sum, 0)
}

func generateCombinations(input, result []int, sum, index int) {
	//Accept case
	if sum == 0 {
		fmt.Printf("\n combination =%v ", result)
		return
	}
	// Reject Case
	if sum < 0 {
		return
	}

	// Steps
	for i := index; i < len(input); i++ {
		generateCombinations(input, append(result, input[i]), sum-input[i], index)
	}
}
