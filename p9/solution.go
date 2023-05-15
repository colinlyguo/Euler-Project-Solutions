package main

import "fmt"

func findPythagoreanTriplet(sum int) (int, int, int) {
	for a := 1; a < sum/3; a++ {
		for b := a + 1; b < sum/2; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return -1, -1, -1
}

func main() {
	a, b, c := findPythagoreanTriplet(1000)
	fmt.Println(a * b * c)
}
