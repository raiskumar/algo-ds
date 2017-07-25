package dp

// (1) Brute-force Approach O(N^3)
// We need to consider all possible subarrays of the array. There will be N(N+1)/2 subarray.
// Then need to calculate sum of each subarray. This will be of order of N

// (2) Non-optimized approach O(N^2)
// Optimized at the cost of O(N) addition space overhead
// Array stored sum till ith index; i.e. Sum[i] = Sum of all elements until index i
// Sum[i,j] = Sum[j] - Sum[i-1]
func maxSubArray(arr []int) (result int) {
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
			}
		}
	}
	return result
}

// (3) DP
// Let say MS(i) is maximum sum ending at index i
// To calcualte for sum for i
// 1) Either add the index to the solution found so far i.e. i-1
// 2) Or, start a new sum from index i
// i.e. MS(i) = Max(MS(i-1)+Arr[i], Arr[i])

func maxSubArrayDp(arr []int) (result int) {
	len := len(arr)
	var sum []int

	sum = append(sum, arr[0])

	for i := 1; i < len; i++ {
		if sum[i-1]+arr[i] > arr[i] {
			sum = append(sum, sum[i-1]+arr[i])
		} else {
			sum = append(sum, arr[i])
		}
	}

	result = sum[0]

	for i := 1; i < len; i++ {
		if result < sum[i] {
			result = sum[i]
		}
	}
	return result
}
