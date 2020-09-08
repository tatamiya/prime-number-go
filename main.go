package main

import (
	"fmt"
	"math"
	"sync"
)

func checkPrimeBasic(target int) bool {
	if target == 2 {
		return true
	}
	if target%2 == 0 {
		return false
	}
	iFrom := 3
	iTo := int(math.Sqrt(float64(target)) + 1)
	for i := iFrom; i < iTo; i += 2 {
		if target%i == 0 {
			return false
		}
	}
	return true
}

var wg sync.WaitGroup

const numProcesses = 4

func checkPrimePool(target int) bool {

	if target == 2 {
		return true
	}
	if target%2 == 0 {
		return false
	}

	iFrom := 3
	iTo := int(math.Sqrt(float64(target)) + 1)
	lengthChunk := int((iTo - iFrom) / numProcesses)

	var isPrime bool
	wg.Add(numProcesses)

	for k := 0; k < numProcesses; k++ {
		iFromChunk := iFrom + lengthChunk*k
		iToChunk := iFromChunk + lengthChunk
		if iToChunk > iTo {
			iToChunk = iTo
		}

		go checkPrimeChunk(target, iFromChunk, iToChunk, &isPrime)
	}
	wg.Wait()
	return isPrime

}

func checkPrimeChunk(target int, iFrom int, iTo int, isPrime *bool) {
	defer wg.Done()
	for i := iFrom; i <= iTo; i += 2 {
		if target%i == 0 {
			return
		}
	}
	*isPrime = true
	return
}

func main() {
	for _, i := range []int{2, 3, 4, 5, 27, 31} {
		isPrime := checkPrimeBasic(i)
		fmt.Printf("Is %d a prime number?: %t\n", i, isPrime)
	}
}
