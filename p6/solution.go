package main

import "fmt"

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func difference(n int) int {
	return squareOfSum(n) - sumOfSquares(n)
}

func main() {
	fmt.Println(difference(100))
}
