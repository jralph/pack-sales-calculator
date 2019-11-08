package calculator

import (
	"sort"
)

func PackCalculator(order int, packs []int) map[int]int {
	sort.Sort(sort.Reverse(sort.IntSlice(packs)))

	smallest := packs[len(packs) - 1]
	divisible := (order + smallest - 1) / smallest
	newOrder := divisible * smallest

	selected := map[int]int{}
	for _, pack := range packs {
		for newOrder >= pack {
			newOrder -= pack
			selected[pack]++
		}
	}

	return selected
}