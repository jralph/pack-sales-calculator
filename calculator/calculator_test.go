package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultPacks = []int{250, 500, 1000, 2000, 5000}
var largePacks = []int{25000, 50000, 100000, 200000, 500000}
var smallPacks = []int{1, 2, 5, 10, 20, 50}

func benchmarkPackCalculator(i int, packs []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = PackCalculator(i, packs)
	}
}

func BenchmarkPackCalculatorNegativeWithDefault(b *testing.B) {
	benchmarkPackCalculator(-1, defaultPacks, b)
}

func BenchmarkPackCalculator0WithDefault(b *testing.B) {
	benchmarkPackCalculator(0, defaultPacks, b)
}

func BenchmarkPackCalculator1WithDefault(b *testing.B) {
	benchmarkPackCalculator(1, defaultPacks, b)
}

func BenchmarkPackCalculator250WithDefault(b *testing.B) {
	benchmarkPackCalculator(250, defaultPacks, b)
}

func BenchmarkPackCalculator500WithDefault(b *testing.B) {
	benchmarkPackCalculator(500, defaultPacks, b)
}

func BenchmarkPackCalculator1000WithDefault(b *testing.B) {
	benchmarkPackCalculator(1000, defaultPacks, b)
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

func BenchmarkPackCalculator250WithAllNegatives(b *testing.B) {
	benchmarkPackCalculator(250, []int{-100, -200}, b)
}

func BenchmarkPackCalculator250WithSomeNegatives(b *testing.B) {
	benchmarkPackCalculator(250, []int{-100, 0, 250, 1000}, b)
}

func TestPackCalculator_returns0For0(t *testing.T) {
	result, err := PackCalculator(0, defaultPacks)

	assert.Empty(t, result)
	assert.Error(t, err, "order amount was negative or zero")
}

func TestPackCalculator_returns0ForNegative(t *testing.T) {
	result, err := PackCalculator(-1, defaultPacks)

	assert.Empty(t, result)
	assert.Error(t, err, "order amount was negative or zero")
}

func TestPackCalculator_returns1x250For1(t *testing.T) {
	result, err := PackCalculator(1, defaultPacks)

	assert.Equal(t, map[int]int{
		250: 1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_returns1x250For250(t *testing.T) {
	result, err := PackCalculator(250, defaultPacks)

	assert.Equal(t, map[int]int{
		250: 1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_returns1x500For251(t *testing.T) {
	result, err := PackCalculator(251, defaultPacks)

	assert.Equal(t, map[int]int{
		500: 1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_returns1x500and1x250For501(t *testing.T) {
	result, err := PackCalculator(501, defaultPacks)

	assert.Equal(t, map[int]int{
		500: 1,
		250: 1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_returns2x2000and1x500and1x250for4750(t *testing.T) {
	result, err := PackCalculator(4750, defaultPacks)

	assert.Equal(t, map[int]int{
		2000: 2,
		500:  1,
		250:  1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_returns2x5000and1x2000and1x250For12001(t *testing.T) {
	result, err := PackCalculator(12001, defaultPacks)

	assert.Equal(t, map[int]int{
		5000: 2,
		2000: 1,
		250:  1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_handlesPacksOfNegatives(t *testing.T) {
	result, err := PackCalculator(1000, []int{-250, -100})

	assert.Empty(t, result)
	assert.Error(t, err, "all packs are negative or zero")
}

func TestPackCalculator_handlesPacksContainingNegatives(t *testing.T) {
	result, err := PackCalculator(250, []int{-100, 250})

	assert.Equal(t, map[int]int{
		250: 1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_handlesPacksWithSomeNegatives(t *testing.T) {
	result, err := PackCalculator(250, []int{-100, -200, 250, 1000})

	assert.Equal(t, map[int]int{
		250: 1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_handlesPacksOfZero(t *testing.T) {
	result, err := PackCalculator(1000, []int{0})

	assert.Empty(t, result)
	assert.Error(t, err, "all packs are negative or zero")
}

func TestPackCalculator_handlesPacksContainingZero(t *testing.T) {
	result, err := PackCalculator(250, []int{0, 250})

	assert.Equal(t, map[int]int{
		250: 1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_handlesNoPacksGiven(t *testing.T) {
	result, err := PackCalculator(250, []int{})

	assert.Empty(t, result)
	assert.Error(t, err, "no packs provided")
}

func TestPackCalculator_handlesPacksThatAreNotMultiples_v1(t *testing.T) {
	result, err := PackCalculator(100, []int{33, 90, 97, 102})

	assert.Equal(t, map[int]int{
		102: 1,
	}, result)
	assert.Nil(t, err)
}

func TestPackCalculator_handlesPacksThatAreNotMultiples_v2(t *testing.T) {
	result, err := PackCalculator(93, []int{33, 90})

	assert.Equal(t, map[int]int{
		33: 3,
	}, result)
	assert.Nil(t, err)
}
