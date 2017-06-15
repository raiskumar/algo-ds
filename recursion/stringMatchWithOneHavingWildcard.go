package recursion

// http://www.geeksforgeeks.org/wildcard-character-matching/
// Given an input string with ? and . as wildcards, check if second string matches input string
// * -> matches 0 or more instances of any character
// . -> maches any one character

func patternMatch(regex, input string) bool {
	if len(regex) == 0 {
		if len(input) == 0 {
			return true
		}
		return false
	}
	if len(input) == 0 {
		return false
	}

	// if the char is * then it will match either 0 or more character of the input
	// Or operator takes care of both cases
	if regex[0] == '*' {
		return patternMatch(regex[1:], input) || // matches 0 char
			patternMatch(regex, input[1:]) // ignore next char of input
	} else if regex[0] == '.' || (regex[0] == input[0]) {
		patternMatch(regex[1:], input[1:])
	} else {
		return false
	}
	return true
}
