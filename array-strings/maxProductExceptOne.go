// Given an array of integers; Find the max product of all elements except one.  Division NOT allowed !
// if arr = {2, 4, 9, 3}   max product = 4*9*3 = 108
// if arr = {-1, -10, 5, 3}  max product = -1*-10*5 = 50

package array

// Brute Force Approach
// Calcualte all possible multiplications for runnig two loolps
// Complexity = O(n^3) ; O(n^2) for loops

// Optimized approach
//The product of all the integers except the integer at each index can be broken down into:
//1. the product of all the integers before each index
//2. the product of all the integers after each index.
// Then multiply at each index and keep updating the max value
// Complexity = O(n); Space Complexity = O(n) for two arrays
func getMaxProductsOfAllIntsExceptOne(arr []int) int {
	if len(arr) < 2 {
		return 0 // invalid case
	}

	var productOfAllIntsBeforeIndex []int
	productSoFar := 1

	//for each integer, find the product of all the integers before it
	for i := 0; i < len(arr); i++ {
		productOfAllIntsBeforeIndex = append(productOfAllIntsBeforeIndex, productSoFar)
		productSoFar *= arr[i]
	}

	productOfAllIntegersAfterIndex := make([]int, len(arr), len(arr))
	productSoFar = 1 // re-initialize

	//for each integer, find the product of all the integers after it
	for i := len(arr) - 1; i >= 0; i-- {
		productOfAllIntegersAfterIndex[i] = productSoFar
		productSoFar *= arr[i]
	}

	maxProduct := 1
	tmpProduct := 1
	// productsOfAllIntsExceptAtIndex = multiplying before and after values
	for i := 0; i < len(arr); i++ {
		tmpProduct = productOfAllIntsBeforeIndex[i] * productOfAllIntegersAfterIndex[i]
		if tmpProduct > maxProduct {
			maxProduct = tmpProduct
		}
	}
	return maxProduct
}

// Above approach uses two temporary arrays; can this be further optized ?
// Can we solve problem with only one temporary array
// The same array can be first used to calcualte product before index and then product after index
// So finally array will have product of all numbers except at the index
func getMaxProductsOfAllIntsExceptOneOptimized(arr []int) int {
	if len(arr) < 2 {
		return 0 // invalid case
	}

	var productOfAllExceptIndex []int

	productSoFar := 1
	for i := 0; i < len(arr); i++ {
		productOfAllExceptIndex = append(productOfAllExceptIndex, productSoFar)
		productSoFar *= arr[i]
	}

	productSoFar = 1 // re-initialize
	for i := len(arr) - 1; i >= 0; i-- {
		productOfAllExceptIndex[i] *= productSoFar
		productSoFar *= arr[i]
	}

	maxProduct := 1
	for i := 0; i < len(arr); i++ {
		if productOfAllExceptIndex[i] > maxProduct {
			maxProduct = productOfAllExceptIndex[i]
		}
	}
	return maxProduct
}
