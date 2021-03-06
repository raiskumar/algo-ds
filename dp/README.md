## Dynamic Programming (DP)
DP is a general technique for solving optimization, search, and computing problmes that can be decomposed into subproblems. 

### Overalpping Subproblem
A problem has overalpping subproblems if finding its solution involves solving the same subproblem multiple times. 

As an example, let's look at fibonacci sequence. 
Recursive implementation :

 ``` 
   fibonacci(n) {
       if n == 0 || n ==1 {
           return n
       }
       febonacci(n-1) + febonacci(n-2)
   }
 ```
 

                  febonacci(n)
                  /           \
              febonacci(n-1)    febonacci(n-2)
              /                     ...
        febonacci(n-2)  
            ...                       


You can clearly note that, fibonacci(n-2) is getting evaluated more than once. Obviously, that's NOT a good thing from performance point of view. Ideally, we would like to perform any operation only once. Let's look how this can be achieved -

### 1) Memoization
Memoization ensures that a function doesn't run for the same inputs more than once by keeping a record of the results for given inputs (usually in a dictionary).
Memoization is a common strategy for dynamic programming problems, which are problems where the solution is composed of solutions to the same problem with smaller inputs (as with the fibonacci problem, above). 

### 2) Going Bottom up
Going bottom-up is a common strategy for dynamic programming problems, which are problems where the solution is composed of solutions to the same problem with smaller inputs (as with the fibonacci problem, above). 

#### DP is quite effective in sovling combinatorial problems, and you can solve DP problems on paper!

## Problems: 
1. [Generate nth Fibonacci number](https://github.com/raiskumar/algo-ds/blob/master/dp/fibonacci.go)
<br /> f(n) = f(n-1) + f(n-2)
<br /> By definition, first two numbers are 0 and 1
<br /> The sequence is 0, 1, 1, 2, 3 ...

2. [Find a subarray of a given array of integers whose sum is MAXIMUM](https://github.com/raiskumar/algo-ds/blob/master/dp/maxSumSubarray.go)
<br /> arr = {904, 40, 523, 12, -335, -385, -124, 481, 31}
<br /> max subarray = arr[0,3] ; both index included
<br /> Hint: Try with brute force approach first and then think of applying DP technique. 

3. [COUNT the number of ways/combinations to reach a given score](https://github.com/raiskumar/algo-ds/blob/master/dp/countScoreCombinations.go)
<br /> All the possible points/runs a player can score are given and also given the total score. Find TOTAL number of ways to reach that given score. 
<br /> Let's say, possible points = {2,3,7}; & total score = 9
<br /> Possible combinations are : {2,2,2,3}, {2,7}, {3,3,3} ; COUNT of combinations = 3
<br /> Please note that, {2,2,2,3} and {2,2,3,2} both are same combination. Think of them as coins so you having 2 coins of denomination 1 and 5 can be expressed as {1,5} or {5,1}
<br /> 


## Testing
 [dp_test.go](dp_test.go) Dynamic Programming implementation of above problems in Golang
