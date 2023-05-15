package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i*i <= n; i += 1 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nthPrime(n int) int {
	counter := 0
	candidate := 1
	for counter < n {
		candidate++
		if isPrime(candidate) {
			counter++
		}
	}
	return candidate
}

func main() {
	fmt.Println(nthPrime(10001))
}
