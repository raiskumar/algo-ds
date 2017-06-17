## Recursion Problems
Here, is reference for learning [Recursion fundamentals](http://geekrai.blogspot.in/2015/08/recursive-thinking.html)

## Backtracking
In backtracking algorithms, we try to build a solution one step at a time. If at some step it becomes clear that the current path cannot lead to a solution, go back to the previous step (backtrack) and choose a different path. Basically once we exhaust all options at a certain, we go back. 

## Exhaustive
Using exhaustive search we consider all subsets irrespective of whether they satisfy given constraints or not. 
At each step, make a choice and then unmake it and try other options until all options are fully exhausted. 


## Problems: 
1. [Write a function to calculate a to the power b, i.e pow(a,b)](https://github.com/raiskumar/algo-ds/blob/master/recursion/power.go)
   <br/>power(2,4) = 2^4 = 16

2. [Find all binary strings that can be formed by replacing wildcard '?' with 0 and 1](https://github.com/raiskumar/algo-ds/blob/master/recursion/allCombinationsOfBinaryStrings.go) 
   <br/>Given a string which can have wildcard '?'. Find all possible strings which can be obtained by replacing ? with 0 and 1.
   <br /> Input = "0?1"
   <br /> Results = "001", "011"   

3. [Generate all binary Strings of length k](https://github.com/raiskumar/algo-ds/blob/master/recursion/generateAllBinaryStrings.go) 
   <br/> Given a length; generate all possible binary strings. 
   <br/> If k = 2; Possible binary strings would be "00", "01", "10", "11"

4. [Generate Power Set of a Set](https://github.com/raiskumar/algo-ds/blob/master/recursion/powerSet.go) 
   <br/> Given a set generate all possible subsets (aka powerset) of the set. 
   <br /> If input is {1,2}; 
   <br /> Output should be {}, {1}, {2}, {1,2}. 

5. [Print all permutations of a String with all unique characters](https://github.com/raiskumar/algo-ds/blob/master/recursion/permutation.go)
    <br/> input = "ab" 
    <br/>  Resuls = "ab", "ba"

6. [Subset Sum problem](https://github.com/raiskumar/algo-ds/blob/master/recursion/subsetSum.go) 
   <br/> Given a set find all subsets (of the array) whose sum adds up to a given number K. Assume that elements are NON-negative. 
   <br /> Input = {1, 2, 9, 3}, Sum = 12
   <br /> Results: {1,2,9}, {9,3}

7. [Combination Sum problem](https://github.com/raiskumar/algo-ds/blob/master/recursion/combinationSum.go)
    <br/> Given a set (without duplicates), find all unique combinations where the candiates add up to a fixed sum. Assume that elements are NON-negative. 
    <br /> This is variation of problem #6 with only difference that the numbers in result can repeat.
    <br /> Input = {2,3,6,7}, Sum = 7
    <br /> Results : {2,2,3}, {2,3,2}, {3,2,2}, {7}

8. [String matching between two string, first one has wildcard characters](https://github.com/raiskumar/algo-ds/blob/master/recursion/stringMatchWithOneHavingWildcard.go) 
   <br/> Given two input string, first can have wildcards * and .; check if second string matches wildcard. Wildcard, * matches any of the 0 or more characters And . matches any single character

9. [String matching between two string, first one has wildcard characters. Variation of above](https://github.com/raiskumar/algo-ds/blob/master/recursion/stringMatchWithOneHavingWildcardV2.go) 
   <br/> Here * can match 0 or more occurances of character just before it

10. [Generate all subsets of size k](https://github.com/raiskumar/algo-ds/blob/master/recursion/kSubsets.go)
   <br/>  A k-subset is a subset of a set on n elements containing exactly k elements. 
   <br/> Input = {1,2,3} 
   <br/> 2 length subsets are {1,2}, {2,3}, & {1,3}

## Testing
 [recursion_test.go](recursion_test.go) Tests implementation of above problems in Golang