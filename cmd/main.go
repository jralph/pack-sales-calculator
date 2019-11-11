package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pack-sales-calculator/calculator"
)

func main() {
	orderString := os.Args[1]
	availablePacks := os.Args[2]

	order, err := strconv.Atoi(orderString)
	if err != nil {
		panic(err)
	}

	packStrings := strings.Split(availablePacks, ",")

	var packs []int
	for _, pack := range packStrings {
		size, err := strconv.Atoi(pack)
		if err != nil {
			panic(err)
		}

		packs = append(packs, size)
	}

	result, err := calculator.PackCalculator(order, packs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}