package array

// Recursive implementation
//   Shifts by one place in each iteration
//   Time Complexity O(nk);
//   Space Complexity O(1)  (Not considering recursion overhead)
func rotate(input []int, k int) {
	if k == 0 {
		return
	}

	tmp := input[len(input)-1]
	for i := len(input) - 1; i > 0; i-- {
		input[i] = input[i-1]
	}
	input[0] = tmp

	rotate(input, k-1)
}

// More optimized approach in O(n) time and O(1) space
//   arr = {1,2,3,4,5,6}
//   Divide the array in 2 parts {1,2,3,4 and 5, 6}
//   1. First part of length n-k
//   2. Second part of length k
//   3. Revers first part {4,3,2,1,5,6}, reverse second part {4,3,2,1,6,5} and then the whole part {5,6,1,2,3,4}
func rotateOptimzed(input []int, k int) {
	if k == 0 {
		return
	}

	a := len(input) - k

	reverse(input, 0, a-1)
	reverse(input, a, len(input)-1)
	reverse(input, 0, len(input)-1)

}

// reverse an array; left and right are indexes of the array in which reversal needs to be performed!
func reverse(input []int, left, right int) {
	tmp := 0
	for left < right {
		tmp = input[left]
		input[left], input[right] = input[right], tmp
		left++
		right--
	}
}
