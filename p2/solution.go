package main

import "fmt"

func sumEvenFibonacci(limit int) int {
	a, b := 1, 2
	total := 0
	for a <= limit {
		if a%2 == 0 {
			total += a
		}
		a, b = b, a+b
	}
	return total
}

func main() {
	fmt.Println(sumEvenFibonacci(4000000))
}
