package main

import "fmt"

func main() {
	maxSeq := 0
	maxNum := 0

	for i := 1; i < 1000000; i++ {
		seq := generateSeq(i)
		if seq > maxSeq {
			maxSeq = seq
			maxNum = i
		}
	}
	fmt.Println(maxNum)
}

func generateSeq(n int) (seq int) {
	for ; n != 1; seq++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return
}
