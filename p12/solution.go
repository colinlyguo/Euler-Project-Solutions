package main

import (
	"fmt"
	"math"
)

// Function to count the divisors of a number
func countDivisors(n int) int {
	count := 0
	root := int(math.Sqrt(float64(n)))
	for i := 1; i <= root; i++ {
		if n%i == 0 {
			if n/i == i {
				// If divisors are equal, count only one
				count++
			} else {
				// Otherwise count both
				count += 2
			}
		}
	}
	return count
}

func main() {
	i := 1
	for {
		// Calculate the ith triangle number
		tri := i * (i + 1) / 2

		// If it has over 500 divisors, print it and exit
		if countDivisors(tri) > 500 {
			fmt.Println(tri)
			break
		}

		i++
	}
}
