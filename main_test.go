package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	largeNonPrime = 100109100129100369
	largePrime    = 100109100129100151
)

func TestCheckPrimeBasic(t *testing.T) {

	assert.True(t, checkPrimeBasic(2))
	assert.True(t, checkPrimeBasic(3))
	assert.False(t, checkPrimeBasic(4))
	assert.True(t, checkPrimeBasic(5))
	assert.False(t, checkPrimeBasic(6))
	assert.True(t, checkPrimeBasic(7))
	assert.False(t, checkPrimeBasic(8))
	assert.False(t, checkPrimeBasic(9))
	assert.False(t, checkPrimeBasic(10))

	assert.False(t, checkPrimeBasic(27))
	assert.True(t, checkPrimeBasic(31))

	assert.False(t, checkPrimeBasic(largeNonPrime))
	assert.True(t, checkPrimeBasic(largePrime))
}

func TestCheckPrimePool(t *testing.T) {

	assert.True(t, checkPrimePool(2))
	assert.True(t, checkPrimePool(3))
	assert.False(t, checkPrimePool(4))
	assert.True(t, checkPrimePool(5))
	assert.False(t, checkPrimePool(6))
	assert.True(t, checkPrimePool(7))
	assert.False(t, checkPrimePool(8))
	assert.False(t, checkPrimePool(9))
	assert.False(t, checkPrimePool(10))

	assert.False(t, checkPrimePool(27))
	assert.True(t, checkPrimePool(31))

	assert.False(t, checkPrimePool(largeNonPrime))
	assert.True(t, checkPrimePool(largePrime))
}

func BenchmarkCheckPrimeCheckBasic_largePrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeBasic(largePrime)
	}
}

func BenchmarkCheckPrimeCheckPool_largePrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimePool(largePrime)
	}
}

func BenchmarkCheckPrimeCheckBasic_largeNonPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimeBasic(largeNonPrime)
	}
}

func BenchmarkCheckPrimeCheckPool_largeNonPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPrimePool(largeNonPrime)
	}
}
