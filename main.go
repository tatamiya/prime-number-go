package main

import (
	"fmt"
	"math"
)

func checkPrimeBasic(target int) bool {
	if target == 2 {
		return true
	}
	if target%2 == 0 {
		return false
	}
	numStart := 3
	numEnd := int(math.Sqrt(float64(target)) + 1)
	for i := numStart; i < numEnd; i += 2 {
		if target%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for _, i := range []int{2, 3, 4, 5, 27, 31} {
		isPrime := checkPrimeBasic(i)
		fmt.Printf("Is %d a prime number?: %t\n", i, isPrime)
	}
}
