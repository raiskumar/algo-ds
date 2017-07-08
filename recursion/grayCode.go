package recursion

// Given a number n, generate bit pattern from 0 to 2^n - 1 such that successive patterns differ by 1 bit

// n-bit Gray Codes can be generated from list of (n-1)-bit Gray codes using following steps.
// 1) Let the list of (n-1)-bit Gray codes be L1. Create another list L2 which is reverse of L1.
// 2) Modify the list L1 by prefixing a ‘0’ in all codes of L1.
// 3) Modify the list L2 by prefixing a ‘1’ in all codes of L2.
// 4) Concatenate L1 and L2. The concatenated list is required list of n-bit Gray codes.
func grayCode(numberOfBits int) (result []string) {
	if numberOfBits == 0 {
		return nil
	}

	result = []string{"0", "1"}
	if numberOfBits == 1 {
		return result
	}

	calculateAllGrayCodes(1, numberOfBits, &result)
	return result
}

func calculateAllGrayCodes(k, n int, result *[]string) {
	if k == n {
		return
	}

	var resultReverse []string
	j := 0
	for i := len(*result) - 1; i >= 0; i-- {
		resultReverse = append(resultReverse, (*result)[i])
		j++
	}

	//Modify the list L1 by prefixing a ‘0’ in all codes of L1.
	for t := 0; t < len(*result); t++ {
		(*result)[t] = "0" + (*result)[t]
	}

	//Modify the list L2 by prefixing a ‘1’ in all codes of L2.
	for q := 0; q < len(resultReverse); q++ {
		resultReverse[q] = "1" + resultReverse[q]
	}

	//append results
	for _, v := range resultReverse {
		(*result) = append((*result), v)
	}

	calculateAllGrayCodes(k+1, n, result)
}
