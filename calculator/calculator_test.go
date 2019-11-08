package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultPacks = []int{250,500,1000,2000,5000}

func TestPackCalculator_returns0For0(t *testing.T) {
	result := PackCalculator(0, defaultPacks)

	assert.Equal(t, result, []int{})
}

func TestPackCalculator_returns1x250For1(t *testing.T) {
	result := PackCalculator(1, defaultPacks)

	assert.Equal(t, result, []int{250})
}

func TestPackCalculator_returns1x250For250(t *testing.T) {
	result := PackCalculator(250, defaultPacks)

	assert.Equal(t, result, []int{250})
}

func TestPackCalculator_returns1x500For251(t *testing.T) {
	result := PackCalculator(251, defaultPacks)

	assert.Equal(t, result, []int{500})
}

func TestPackCalculator_returns1x500and1x250For501(t *testing.T) {
	result := PackCalculator(501, defaultPacks)

	assert.Equal(t, result, []int{500, 250})
}

func TestPackCalculator_returns2x2000and1x500and1x250for4750(t *testing.T) {
	result := PackCalculator(4750, defaultPacks)

	assert.Equal(t, result, []int{2000, 2000, 500, 250})
}

func TestPackCalculator_returns2x5000and1x2000and1x250For12001(t *testing.T) {
	result := PackCalculator(12001, defaultPacks)

	assert.Equal(t, result, []int{5000, 5000, 2000, 250})
}