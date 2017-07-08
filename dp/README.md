## Problems related with Dynamic Programming (DP)
DP is a general technique for solving optimization, search, and computing problmes that can be decomposed into subproblems. 

#Overalpping Subproblem
A problem has overalpping subproblems if finding its solution involves solving the same subproblem multiple times. 

As an example, let's look at fibonacci sequence. 
Recursive implementation :
   febonacci(n) {
       if n == 0 || n ==1 {
           return n
       }
       febonacci(n-1) + febonacci(n-2)
   }

                  febonacci(k)
                  /           \
              febonacci(k-1)    febonacci(k-2)
              /                     ...
        febonacci(k-2)  
            ...                       

    You can clearly note that, fibonacci(k-2) is getting evaluated more than once. 

 # To get rid of above problem, a technique known as Memoization is used
Memoization ensures that a function doesn't run for the same inputs more than once by keeping a record of the results for given inputs (usually in a dictionary).
Memoization is a common strategy for dynamic programming problems, which are problems where the solution is composed of solutions to the same problem with smaller inputs (as with the fibonacci problem, above). 

# Another approach is going Bottom up
Going bottom-up is a common strategy for dynamic programming problems, which are problems where the solution is composed of solutions to the same problem with smaller inputs (as with the fibonacci problem, above). 

## Problems: 
1. [Fibonacci Sequence](https://github.com/raiskumar/algo-ds/blob/master/dp/power.go)
   <br/>power(2,4) = 2^4 = 16

## Testing
 [recursion_test.go](arrays_test.go) Tests implementation of above problems in Golang
