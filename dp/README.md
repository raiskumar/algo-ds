## Dynamic Programming (DP)
DP is a general technique for solving optimization, search, and computing problmes that can be decomposed into subproblems. 

### Overalpping Subproblem
A problem has overalpping subproblems if finding its solution involves solving the same subproblem multiple times. 

As an example, let's look at fibonacci sequence. 

Recursive implementation :

 ```febonacci(n) {
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

## Problems: 
1. [Fibonacci Sequence](https://github.com/raiskumar/algo-ds/blob/master/dp/fibonacci.go)

## Testing
 [recursion_test.go](dp_test.go) Dynamic Programming implementation of above problems in Golang
