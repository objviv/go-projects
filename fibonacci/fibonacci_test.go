package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciLessThanZero(t *testing.T) {
	assert.Panics(t, func() { fibonacci(-1) })
}

func TestFibonacciZero(t *testing.T) {
	ret := fibonacci(0)
	if ret != 0 {
		t.Errorf("This should be zero")
	}
}

func TestFibonacciTwo(t *testing.T) {
	ret := fibonacci(2)
	if ret != 1 {
		t.Errorf("This should be 1")
	}
}

func TestFibonacciFive(t *testing.T) {
	ret := fibonacci(5)
	if ret != 5 {
		t.Errorf("This should be 5")
	}
}

func TestFibonacciTen(t *testing.T) {
	ret := fibonacci(10)
	if ret != 55 {
		t.Errorf("This should be 55")
	}
}
