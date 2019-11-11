package calculator

import (
	"errors"
	"sort"
)

// PackCalculator calculates the minimum amount of items to
// send to a user, while also sending out as few packs as possible
// and never using partial packs.
func PackCalculator(orderAmount int, packs []int) (map[int]int, error) {
	selected := make(map[int]int)

	// If this is 0, we just return an empty selection of packs and an error.
	// This is about 4x more efficient than letting the order
	// for 0 proceed through the rest of the algorithm.
	if orderAmount <= 0 {
		return selected, errors.New("order amount was negative or zero")
	}

	// If we have no packs, return an empty selection and error.
	if len(packs) == 0 {
		return selected, errors.New("no packs provided")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(packs)))

	// Remove any negative packs as these are invalid.
	for i := 0; i < len(packs); i++ {
		if packs[i] <= 0 {
			packs = append(packs[:i], packs[i+1:]...)
			i--
		}
	}

	// If we now have no packs, that is because they are all negative
	// or zero.
	if len(packs) == 0 {
		return selected, errors.New("all packs are negative or zero")
	}

	smallest := packs[len(packs) - 1]

	// If the order is equal to or smaller than the smallest pack
	// we return early. This is slightly more efficient.
	if orderAmount <= smallest {
		selected[smallest]++
		return selected, nil
	}

	// If the order is equal to any of the packs we have we
	// return early. This is slightly more efficient.
	for _, pack := range packs {
		if orderAmount == pack {
			selected[pack]++
			return selected, nil
		}
	}

	// Calculate the orderAmount rounded up to the
	// closest number of the smallest pack that fit into it.
	// eg, 501 becomes 750 if the smallest pack is 250
	timesDivisible := (orderAmount + smallest - 1) / smallest
	newOrderAmount := timesDivisible * smallest

	for _, pack := range packs {
		for newOrderAmount >= pack {
			newOrderAmount -= pack
			selected[pack]++
		}
	}

	return selected, nil
}