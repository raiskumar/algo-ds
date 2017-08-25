package array

// Given an array; find the index of longest contiguous subarray
// {2,11,3,5,13,7,19,17,233} result = (2,4) i.e. array starting at index 2 and ending at index 4 {3,5,13}

// Brute force approach - consider all possible subarray and then check; the one with more element will be result.
// Not an optimized approach

// Optimized
// We can do one iteration and keep track of maximum length increasing subarray seen so far
func longestContiguousIncreasingSubarray(arr []int) (start int, end int) {
	if len(arr) < 2 {
		return 0, len(arr) - 1
	}

	size := 1             // size of subarray
	index := -1           // starting index of subarray
	maxSizeSeenSoFar := 1 // maximum length of countiguous increasing subarry seen so far
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] { // if increasing subarray
			size++
			if index == -1 {
				index = i - 1
			}
		} else {
			size = 1
			index = -1
		}
		if maxSizeSeenSoFar < size { // reset if current value is less than previous value
			maxSizeSeenSoFar = size
			start = index
			end = maxSizeSeenSoFar - 1
		}

	}
	return start, end
}
