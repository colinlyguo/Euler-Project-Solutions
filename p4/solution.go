package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func largestPalindromeProduct() int {
	max := 0
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if max >= product {
				break
			}
			if isPalindrome(product) {
				max = product
			}
		}
	}
	return max
}

func main() {
	fmt.Println(largestPalindromeProduct())
}
