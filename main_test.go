package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	largeNonPrime = 100109100129100369
	largePrime    = 100109100129100151
)

func TestCheckPrimeSequential(t *testing.T) {

	assert.True(t, checkPrimeSequential(2))
	assert.True(t, checkPrimeSequential(3))
	assert.False(t, checkPrimeSequential(4))
	assert.True(t, checkPrimeSequential(5))
	assert.False(t, checkPrimeSequential(6))
	assert.True(t, checkPrimeSequential(7))
	assert.False(t, checkPrimeSequential(8))
	assert.False(t, checkPrimeSequential(9))
	assert.False(t, checkPrimeSequential(10))

	assert.False(t, checkPrimeSequential(27))
	assert.True(t, checkPrimeSequential(31))

	assert.False(t, checkPrimeSequential(largeNonPrime))
	assert.True(t, checkPrimeSequential(largePrime))
}

//func TestCheckPrimeParallel(t *testing.T) {
//
//	assert.True(t, checkPrimeParallel(2))
//	assert.True(t, checkPrimeParallel(3))
//	assert.False(t, checkPrimeParallel(4))
//	assert.True(t, checkPrimeParallel(5))
//	assert.False(t, checkPrimeParallel(6))
//	assert.True(t, checkPrimeParallel(7))
//	assert.False(t, checkPrimeParallel(8))
//	assert.False(t, checkPrimeParallel(9))
//	assert.False(t, checkPrimeParallel(10))
//
//	assert.False(t, checkPrimeParallel(27))
//	assert.True(t, checkPrimeParallel(31))
//
//	assert.False(t, checkPrimeParallel(largeNonPrime))
//	assert.True(t, checkPrimeParallel(largePrime))
//}

func TestCheckPrimeIPC(t *testing.T) {

	assert.True(t, checkPrimeIPC(2))
	assert.True(t, checkPrimeIPC(3))
	assert.False(t, checkPrimeIPC(4))
	assert.True(t, checkPrimeIPC(5))
	assert.False(t, checkPrimeIPC(6))
	assert.True(t, checkPrimeIPC(7))
	assert.False(t, checkPrimeIPC(8))
	assert.False(t, checkPrimeIPC(9))
	assert.False(t, checkPrimeIPC(10))

	assert.False(t, checkPrimeIPC(27))
	assert.True(t, checkPrimeIPC(31))

	assert.False(t, checkPrimeIPC(largeNonPrime))
	assert.True(t, checkPrimeIPC(largePrime))
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
