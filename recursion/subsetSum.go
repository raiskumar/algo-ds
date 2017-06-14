package recursion

import "fmt"

// Subset sum problem is to find subset of elements that are selected from a given set whose sum adds up to a given number K.
// Assume that the set contains non-negative values.
// It is assumed that the input set is unique (no duplicates are presented).

func subsetSum(arr []int, sum int) {
	var output []int
	findSubArrays(arr, output, sum, 0)
}

// output array stores the subset of the input
// index is the current index of input under consideration
func findSubArrays(input, output []int, sum, index int) {
	//fmt.Println("output =%v, %d", output, sum)

	// Accept (success) case
	if sum == 0 {
		fmt.Printf("\n Subset =%v", output)
		return
	}

	// Reject Case
	// This condition will help in backtracking to the previous step
	if sum < 0 {
		len := len(output)
		sum = sum + output[len-1]
		return
	}

	//Steps
	for i := index; i < len(input); i++ {
		output = append(output, input[i]) // append the new value from the input to output
		index++
		findSubArrays(input, output, sum-input[i], index)

		//want to try out all possible options; so the added input element should be removed from the output
		// sum doesn't need to be updated as it's value was only modified for the function call
		l := len(output)
		output = output[0 : l-1]
	}

}
