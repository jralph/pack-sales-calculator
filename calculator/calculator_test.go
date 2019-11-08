package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultPacks = []int{250,500,1000,2000,5000}
var largePacks = []int{25000, 50000, 100000, 200000, 500000}
var smallPacks = []int{1, 2, 5, 10, 20, 50}

func benchmarkPackCalculator(i int, packs []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = PackCalculator(i, packs)
	}
}

func BenchmarkPackCalculator0WithDefault(b *testing.B) {
	benchmarkPackCalculator(0, defaultPacks, b)
}

func BenchmarkPackCalculator12001WithDefault(b *testing.B) {
	benchmarkPackCalculator(12001, defaultPacks, b)
}

func BenchmarkPackCalculator12001WithLarge(b *testing.B) {
	benchmarkPackCalculator(12001, largePacks, b)
}

func BenchmarkPackCalculator47501043056WithDefault(b *testing.B) {
	benchmarkPackCalculator(47501043056, defaultPacks, b)
}

func BenchmarkPackCalculator47501043056WithLarge(b *testing.B) {
	benchmarkPackCalculator(47501043056, largePacks, b)
}

func BenchmarkPackCalculator12001WithSmall(b *testing.B) {
	benchmarkPackCalculator(12001, smallPacks, b)
}

func TestPackCalculator_returns0For0(t *testing.T) {
	result := PackCalculator(0, defaultPacks)

	assert.Empty(t, result)
}

func TestPackCalculator_returns1x250For1(t *testing.T) {
	result := PackCalculator(1, defaultPacks)

	assert.Equal(t, result, map[int]int{
		250: 1,
	})
}

func TestPackCalculator_returns1x250For250(t *testing.T) {
	result := PackCalculator(250, defaultPacks)

	assert.Equal(t, result, map[int]int{
		250: 1,
	})
}

func TestPackCalculator_returns1x500For251(t *testing.T) {
	result := PackCalculator(251, defaultPacks)

	assert.Equal(t, result, map[int]int{
		500: 1,
	})
}

func TestPackCalculator_returns1x500and1x250For501(t *testing.T) {
	result := PackCalculator(501, defaultPacks)

	assert.Equal(t, result, map[int]int{
		500: 1,
		250: 1,
	})
}

func TestPackCalculator_returns2x2000and1x500and1x250for4750(t *testing.T) {
	result := PackCalculator(4750, defaultPacks)

	assert.Equal(t, result, map[int]int{
		2000: 2,
		500: 1,
		250: 1,
	})
}

func TestPackCalculator_returns2x5000and1x2000and1x250For12001(t *testing.T) {
	result := PackCalculator(12001, defaultPacks)

	assert.Equal(t, result, map[int]int{
		5000: 2,
		2000: 1,
		250: 1,
	})
}