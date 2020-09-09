package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	largeNonPrime1 = 100109100129100369
	largeNonPrime2 = 100109100129101027
	largePrime1    = 100109100129100151
	largePrime2    = 100109100129162907
)

func testHelper(t *testing.T, checkPrime func(int) bool) {

	assert.True(t, checkPrime(2))
	assert.True(t, checkPrime(3))
	assert.False(t, checkPrime(4))
	assert.True(t, checkPrime(5))
	assert.False(t, checkPrime(6))
	assert.True(t, checkPrime(7))
	assert.False(t, checkPrime(8))
	assert.False(t, checkPrime(9))
	assert.False(t, checkPrime(10))

	assert.False(t, checkPrime(27))
	assert.True(t, checkPrime(31))

	assert.False(t, checkPrime(largeNonPrime1))
	assert.False(t, checkPrime(largeNonPrime2))
	assert.True(t, checkPrime(largePrime1))
	assert.True(t, checkPrime(largePrime2))
}

func TestCheckPrimeSequential(t *testing.T) {
	testHelper(t, checkPrimeSequential)
}

func TestCheckPrimeIPC(t *testing.T) {
	testHelper(t, checkPrimeIPC)
}

func TestCheckPrimeParallel(t *testing.T) {
	testHelper(t, checkPrimeParallel)
}

func BenchmarkCheckPrimeSequential_largePrime1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeSequential(largePrime1)
	}
}

func BenchmarkCheckPrimeIPC_largePrime1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeIPC(largePrime1)
	}
}

func BenchmarkCheckPrimeSequential_largeNonPrime1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeSequential(largeNonPrime1)
	}
}

func BenchmarkCheckPrimeIPC_largeNonPrime1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeIPC(largeNonPrime1)
	}
}

func BenchmarkCheckPrimeSequential_largePrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeSequential(largePrime2)
	}
}

func BenchmarkCheckPrimeIPC_largePrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeIPC(largePrime2)
	}
}

func BenchmarkCheckPrimeSequential_largeNonPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeSequential(largeNonPrime2)
	}
}

func BenchmarkCheckPrimeIPC_largeNonPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeIPC(largeNonPrime2)
	}
}
