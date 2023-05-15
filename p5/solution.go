package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

func smallestMultiple(n int) int {
	multiple := 1
	for i := 2; i <= n; i++ {
		multiple = lcm(multiple, i)
	}
	return multiple
}

func main() {
	fmt.Println(smallestMultiple(20))
}
