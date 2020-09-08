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

	// If iTo - iFrom = 13 and numProcesses = 4
	// then, lengthChunk = 4 -> chunked into (4, 4, 4, 1)
	// Without +1, lengthChunk = 3 -> chunked into (3, 3, 3, 3), leaving 13rd
	lengthChunk := int((iTo-iFrom)/numProcesses) + 1

	var isPrime bool = true
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
	if iFrom > iTo {
		return
	}

	for i := iFrom; i < iTo; i += 2 {
		if target%i == 0 {
			*isPrime = false
			return
		}
	}
	return
}

func main() {
	for _, i := range []int{2, 3, 4, 5, 27, 31} {
		isPrime := checkPrimeBasic(i)
		fmt.Printf("Is %d a prime number?: %t\n", i, isPrime)
	}
}
