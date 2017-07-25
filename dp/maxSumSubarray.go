package dp

// (1) Brute-force Approach O(N^3)
// We need to consider all possible subarrays of the array. There will be N(N+1)/2 subarray.
// Then need to calculate sum of each subarray. This will be of order of N

// (2) Non-optimized approach O(N^2)
// Optimized at the cost of O(N) addition space overhead
// Array stored sum till ith index; i.e. Sum[i] = Sum of all elements until index i
// Sum[i,j] = Sum[j] - Sum[i-1]

// (3) DP
//
