package main

import "fmt"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println("for testing ")
}
