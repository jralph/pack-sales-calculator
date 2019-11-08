package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gymshark-test/calculator"
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

	fmt.Println(calculator.PackCalculator(order, packs))
}