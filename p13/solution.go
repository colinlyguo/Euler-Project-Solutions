package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []string{}
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	sum := new(big.Int)
	for _, number := range numbers {
		// 只考虑每个数的前15位
		tmp := new(big.Int)
		tmp.SetString(number[:15], 10)
		sum.Add(sum, tmp)
	}

	fmt.Println(sum.String()[:10])
}
