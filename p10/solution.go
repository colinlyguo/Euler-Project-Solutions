package main

import (
	"fmt"
)

func sumPrimes(n int) int64 {
	sieve := make([]bool, n)
	for i := 2; i < n; i++ {
		sieve[i] = true
	}

	for p := 2; p*p < n; p++ {
		if sieve[p] {
			for i := p * p; i < n; i += p {
				sieve[i] = false
			}
		}
	}

	var sum int64
	for i := 2; i < n; i++ {
		if sieve[i] {
			sum += int64(i)
		}
	}

	return sum
}

func main() {
	fmt.Println(sumPrimes(2000000))
}
