// Given a binary pattern that contains ? as wildcard character. Find all combinates of string that can be formed
// by replacing the wildchard character with either 0 or 1
// input : "?1"
// Output : "01", "11"

package recursion

import "fmt"
import "strings"

func allBinaryStrings(input string) {
	if len(input) == 0 {
		return
	}

	recurse(input, 0)
}

// Process each character one by one and recurse the remaining one
// If the char is either 0 or 1 then ignore it and move to next character
// If it's '?' then replace it with 0 and 1 and recurse to next character
// It will be done in DFS manner
func recurse(input string, index int) {
	if index >= len(input) {
		if !strings.Contains(input, "?") {
			fmt.Println(input)
		}
		return
	}

	if input[index] != '?' {
		recurse(input, index+1)
	} else {
		if index != len(input)-1 {
			recurse(input[0:index]+"0"+input[index+1:], index+1)
			recurse(input[0:index]+"1"+input[index+1:], index+1)
		} else {
			recurse(input[0:index]+"0", index+1)
			recurse(input[0:index]+"1", index+1)
		}
	}
}
