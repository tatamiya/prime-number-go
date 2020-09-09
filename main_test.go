package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	largeNonPrime = 100109100129100369
	largePrime    = 100109100129100151
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

	assert.False(t, checkPrime(largeNonPrime))
	assert.True(t, checkPrime(largePrime))
}

func TestCheckPrimeSequential(t *testing.T) {
	testHelper(t, checkPrimeSequential)
}

func TestCheckPrimeIPC(t *testing.T) {
	testHelper(t, checkPrimeIPC)
}

func BenchmarkCheckPrimeSequential_largePrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeSequential(largePrime)
	}
}

func BenchmarkCheckPrimeIPC_largePrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeIPC(largePrime)
	}
}

func BenchmarkCheckPrimeSequential_largeNonPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeSequential(largeNonPrime)
	}
}

func BenchmarkCheckPrimeIPC_largeNonPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeIPC(largeNonPrime)
	}
}
