## Recursion and Backtracking related Problems

## Backtracking

## Exhaustive
Using exhaustive search we consider all subsets irrespective of whether they satisfy given constraints or not. 


## Problems: 
- [Calculate Power; pow(a,b)](https://github.com/raiskumar/algo-ds/blob/master/recursion/power.go) : pow(2,4) = 16
- [Find all binary strings that can be formed by replacing wildcard '?' with 0 and 1](https://github.com/raiskumar/algo-ds/blob/master/recursion/allCombinationsOfBinaryStrings.go) : Given a string which can have wildcard '?'. Find all possible strings which can be obtained by replacing ? with 0 and 1. 
- [Generate all binary Strings of length k](https://github.com/raiskumar/algo-ds/blob/master/recursion/generateAllBinaryStrings.go) : Given a length; generate all possible binary strings. If k = 2; Possible binary strings would be "00", "01", "10", "11"
- [Generate Power Set of a Set](https://github.com/raiskumar/algo-ds/blob/master/recursion/powerSet.go) : Given a set generate all possible subsets (aka powerset) of the set. If input is {1,2}; Output should be {}, {1}, {2}, {1,2}. 
- [Print all permutations of a String with all unique characters](https://github.com/raiskumar/algo-ds/blob/master/recursion/permutation.go)
- [Subset Sum problem](https://github.com/raiskumar/algo-ds/blob/master/recursion/subsetSum.go) : Given an array find all subsets (of the array) whose sum adds up to a given number K
- [String matching between two string, first one has wildcard characters](https://github.com/raiskumar/algo-ds/blob/master/recursion/stringMatchWithOneHavingWildcard.go) : Given two input string, first can have wildcards * and .; check if second string matches wildcard. Wildcard, * matches any of the 0 or more characters And . matches any single character
- [String matching between two string, first one has wildcard characters. Variation of above](https://github.com/raiskumar/algo-ds/blob/master/recursion/stringMatchWithOneHavingWildcardV2.go) : Here * can match 0 or more occurances of character just before it

## Testing
 [recursion_test.go](recursion_test.go) can be used to test any implementation.

