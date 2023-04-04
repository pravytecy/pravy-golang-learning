package main

import (
	"fmt"
	"math"
)

func main() {
	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(printPrime(input))
}

func printEven(input []int) []int {
	i := 0
	output := []int{}
	for ; i < len(input); i++ {
		if input[i]%2 == 0 {
			output = append(output, input[i])
		}
	}
	return output
}
func printOdd(input []int) []int {
	i := 0
	output := []int{}
	for ; i < len(input); i++ {
		if input[i]%2 != 0 {
			output = append(output, input[i])
		}
	}
	return output
}

func printPrime(input []int) []int {
	i := 0
	output := []int{}
	for ; i < len(input); i++ {

		if isPrime(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
