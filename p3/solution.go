package main

import "fmt"

func largestPrimeFactor(n int64) int64 {
	var i, maxPrime int64
	maxPrime = -1

	for i = 2; i*i <= n; i += 1 {
		for n%i == 0 {
			maxPrime = i
			n /= i
		}
	}

	if n > maxPrime {
		maxPrime = n
	}

	return maxPrime
}

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}
