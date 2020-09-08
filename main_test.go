package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.True(t, checkPrimeBasic(30011))
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

	assert.False(t, checkPrimeBasic(27))
	assert.True(t, checkPrimeBasic(31))

	// ToDo: Bench mark test for large prime / non-prime numbers
	assert.True(t, checkPrimeBasic(30011))
}
