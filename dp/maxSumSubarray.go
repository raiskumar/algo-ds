package dp

// (1) Brute-force Approach O(N^3)
// We need to consider all possible subarrays of the array. There will be N(N+1)/2 subarray.
// Then need to calculate sum of each subarray. This will be of order of N

// (2) Non-optimized approach O(N^2)
// Optimized at the cost of O(N) addition space overhead
// Array stored sum till ith index; i.e. Sum[i] = Sum of all elements until index i
// Sum[i,j] = Sum[j] - Sum[i-1]
func maxSubArray(arr []int) (start, end, result int) {
	len := len(arr)
	var sum []int

	// populate sum array
	sum = append(sum, arr[0])
	for i := 1; i < len; i++ {
		sum = append(sum, sum[i-1]+arr[i]) //sum[i] = sum[i-1] + arr[i]
	}

	sumjk := -1
	for j := 0; j < len-1; j++ {
		for k := 1; k < len; k++ {

			if j-1 == -1 {
				sumjk = sum[k] // sum[j-1] = 0
			} else {
				sumjk = sum[k] - sum[j-1] //sum(j,k) = sum(k) - sum(j - 1)
			}

			// Update result variables
			if result < sumjk {
				result = sumjk
				start = j
				end = k
			}
		}
	}
	return start, end, result
}

// (3) DP
//
