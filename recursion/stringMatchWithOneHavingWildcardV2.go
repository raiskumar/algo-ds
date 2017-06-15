package recursion

// http://www.geeksforgeeks.org/wildcard-character-matching/
// Given an input string with ? and . as wildcards, check if second string matches input string
// * -> matches 0 or more occurances of character just before it
// . -> maches any one character

func patternMatchV2(regex, input string) bool {
	if len(regex) == 0 {
		if len(input) == 0 {
			return true
		}
		return false
	}
	if len(input) == 0 {
		return false
	}

	if len(regex) > 1 && regex[1] == '*' {
		if input[0] == regex[0] {
			return patternMatch(regex, input[1:]) || // matches 0 char
				patternMatch(regex[2:], input[1:]) // ignore next char of input
		}
		return false
	} else if regex[0] == '.' || (regex[0] == input[0]) {
		patternMatch(regex[1:], input[1:])
	} else {
		return false
	}
	return true
}
